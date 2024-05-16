package mappingInfra

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/entity"
	"github.com/speedianet/os/src/domain/valueObject"
	internalDbInfra "github.com/speedianet/os/src/infra/internalDatabase"
	dbModel "github.com/speedianet/os/src/infra/internalDatabase/model"
)

type MappingQueryRepo struct {
	persistentDbSvc *internalDbInfra.PersistentDatabaseService
}

func NewMappingQueryRepo(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
) *MappingQueryRepo {
	return &MappingQueryRepo{
		persistentDbSvc: persistentDbSvc,
	}
}

func (repo *MappingQueryRepo) ReadById(
	id valueObject.MappingId,
) (entity entity.Mapping, err error) {
	model := dbModel.Mapping{}
	err = repo.persistentDbSvc.Handler.
		Model(&dbModel.Mapping{}).
		Where("ID = ?", id.Get()).
		First(&model).Error
	if err != nil {
		return entity, errors.New("ReadDatabaseEntryError")
	}

	entity, err = model.ToEntity()
	if err != nil {
		return entity, errors.New("ModelToEntityError")
	}

	return entity, nil
}

func (repo *MappingQueryRepo) ReadByHostname(
	hostname valueObject.Fqdn,
) (entities []entity.Mapping, err error) {
	models := []dbModel.Mapping{}
	err = repo.persistentDbSvc.Handler.
		Model(&dbModel.Mapping{}).
		Where("Hostname = ?", hostname.String()).
		Find(&models).Error
	if err != nil {
		return entities, errors.New("ReadDatabaseEntriesError")
	}

	for _, model := range models {
		entity, err := model.ToEntity()
		if err != nil {
			log.Printf("ModelToEntityError: %s", err.Error())
			continue
		}

		entities = append(entities, entity)
	}

	return entities, nil
}

func (repo *MappingQueryRepo) ReadByServiceName(
	serviceName valueObject.ServiceName,
) (entities []entity.Mapping, err error) {
	models := []dbModel.Mapping{}
	err = repo.persistentDbSvc.Handler.
		Model(&dbModel.Mapping{}).
		Where("TargetType = service AND TargetValue = ?", serviceName.String()).
		Find(&models).Error
	if err != nil {
		return entities, errors.New("ReadDatabaseEntriesError")
	}

	for _, model := range models {
		entity, err := model.ToEntity()
		if err != nil {
			log.Printf("ModelToEntityError: %s", err.Error())
			continue
		}

		entities = append(entities, entity)
	}

	return entities, nil
}
