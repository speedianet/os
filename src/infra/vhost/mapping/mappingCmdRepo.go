package mappingInfra

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/entity"
	"github.com/speedianet/os/src/domain/valueObject"
	infraHelper "github.com/speedianet/os/src/infra/helper"
	internalDbInfra "github.com/speedianet/os/src/infra/internalDatabase"
	dbModel "github.com/speedianet/os/src/infra/internalDatabase/model"
	vhostInfra "github.com/speedianet/os/src/infra/vhost"
)

type MappingCmdRepo struct {
	persistentDbSvc  *internalDbInfra.PersistentDatabaseService
	mappingQueryRepo *MappingQueryRepo
}

func NewMappingCmdRepo(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
) *MappingCmdRepo {
	mappingQueryRepo := NewMappingQueryRepo(persistentDbSvc)

	return &MappingCmdRepo{
		persistentDbSvc:  persistentDbSvc,
		mappingQueryRepo: mappingQueryRepo,
	}
}

func (repo *MappingCmdRepo) mappingConfigFactory(
	mapping entity.Mapping,
) (string, error) {
	nginxConfig := "location " + mapping.Path.String() + " {"

	switch mapping.TargetType.String() {
	case "url":
		nginxConfig += `
	return 301 ` + mapping.TargetUrl.String() + `;`
	case "service":
		nginxConfig += ``
	case "response-code":
		nginxConfig += `
	return ` + mapping.TargetHttpResponseCode.String() + `;`
	case "inline-html":
		nginxConfig += `
	add_header Content-Type text/html;
	return 200 ` + mapping.TargetInlineHtmlContent.String() + `;`
	case "static-files":
		nginxConfig += `
	try_files $uri $uri/ index.html?$query_string;`
	}

	nginxConfig += `
}
`
	return nginxConfig, nil
}

func (repo *MappingCmdRepo) rebuildMappingFile(
	mappingHostname valueObject.Fqdn,
) error {
	mappings, err := repo.mappingQueryRepo.GetByHostname(mappingHostname)
	if err != nil {
		return err
	}

	vhostQueryRepo := vhostInfra.VirtualHostQueryRepo{}
	mappingFilePath, err := vhostQueryRepo.GetVirtualHostMappingsFilePath(
		mappingHostname,
	)
	if err != nil {
		return errors.New("GetVirtualHostMappingsFilePathError: " + err.Error())
	}

	fullMappingConfigContent := ""
	for _, mapping := range mappings {
		mappingConfigContent, err := repo.mappingConfigFactory(mapping)
		if err != nil {
			log.Printf(
				"MappingConfigFactoryError (%s): %s",
				mapping.Path.String(),
				err.Error(),
			)
		}
		fullMappingConfigContent += mappingConfigContent
	}

	shouldOverwrite := true
	return infraHelper.UpdateFile(
		mappingFilePath.String(),
		fullMappingConfigContent,
		shouldOverwrite,
	)
}

func (repo *MappingCmdRepo) Create(
	createDto dto.CreateMapping,
) (valueObject.MappingId, error) {
	var mappingId valueObject.MappingId

	vhostCmdRepo := vhostInfra.VirtualHostCmdRepo{}

	isServiceMapping := createDto.TargetType.String() == "service"
	isPhpServiceMapping := isServiceMapping && createDto.TargetServiceName.String() == "php"
	if isPhpServiceMapping {
		err := vhostCmdRepo.CreatePhpVirtualHost(createDto.Hostname)
		if err != nil {
			return mappingId, err
		}
	}

	mappingModel := dbModel.Mapping{}.AddDtoToModel(createDto)
	createResult := repo.persistentDbSvc.Handler.Create(&mappingModel)
	if createResult.Error != nil {
		return mappingId, createResult.Error
	}
	mappingId, err := valueObject.NewMappingId(mappingModel.ID)
	if err != nil {
		return mappingId, err
	}

	err = repo.rebuildMappingFile(createDto.Hostname)
	if err != nil {
		return mappingId, err
	}

	return mappingId, vhostCmdRepo.ReloadWebServer()
}
