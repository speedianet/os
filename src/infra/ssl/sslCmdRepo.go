package sslInfra

import (
	"errors"
	"log"
	"os"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/entity"
	"github.com/speedianet/os/src/domain/valueObject"
	infraHelper "github.com/speedianet/os/src/infra/helper"
	infraData "github.com/speedianet/os/src/infra/infraData"
	internalDbInfra "github.com/speedianet/os/src/infra/internalDatabase"
	o11yInfra "github.com/speedianet/os/src/infra/o11y"
	vhostInfra "github.com/speedianet/os/src/infra/vhost"
	mappingInfra "github.com/speedianet/os/src/infra/vhost/mapping"
)

type SslCmdRepo struct {
	persistentDbSvc *internalDbInfra.PersistentDatabaseService
	transientDbSvc  *internalDbInfra.TransientDatabaseService
	sslQueryRepo    SslQueryRepo
}

func NewSslCmdRepo(
	persistentDbSvc *internalDbInfra.PersistentDatabaseService,
	transientDbSvc *internalDbInfra.TransientDatabaseService,
) *SslCmdRepo {
	return &SslCmdRepo{
		persistentDbSvc: persistentDbSvc,
		transientDbSvc:  transientDbSvc,
		sslQueryRepo:    SslQueryRepo{},
	}
}

func (repo *SslCmdRepo) deleteCurrentSsl(vhost valueObject.Fqdn) error {
	vhostStr := vhost.String()

	vhostCertFilePath := infraData.GlobalConfigs.PkiConfDir + "/" + vhostStr + ".crt"
	vhostCertFileExists := infraHelper.FileExists(vhostCertFilePath)
	if vhostCertFileExists {
		err := os.Remove(vhostCertFilePath)
		if err != nil {
			return errors.New("DeleteCertFileError: " + err.Error())
		}
	}

	vhostCertKeyFilePath := infraData.GlobalConfigs.PkiConfDir + "/" + vhostStr + ".key"
	vhostCertKeyFileExists := infraHelper.FileExists(vhostCertKeyFilePath)
	if vhostCertKeyFileExists {
		err := os.Remove(vhostCertKeyFilePath)
		if err != nil {
			return errors.New("DeleteCertKeyFileError: " + err.Error())
		}
	}

	return nil
}

func (repo *SslCmdRepo) ReplaceWithSelfSigned(vhostName valueObject.Fqdn) error {
	vhostQueryRepo := vhostInfra.NewVirtualHostQueryRepo(repo.persistentDbSvc)
	aliases, err := vhostQueryRepo.ReadAliasesByParentHostname(vhostName)
	if err != nil {
		return errors.New("ReadVhostAliasesError: " + err.Error())
	}

	aliasesHostname := []string{}
	for _, alias := range aliases {
		aliasesHostname = append(aliasesHostname, alias.Hostname.String())
	}

	err = repo.deleteCurrentSsl(vhostName)
	if err != nil {
		return err
	}

	return infraHelper.CreateSelfSignedSsl(
		infraData.GlobalConfigs.PkiConfDir,
		vhostName.String(),
		aliasesHostname,
	)
}

func (repo *SslCmdRepo) dnsFilterFunctionalHostnames(
	vhostNames []valueObject.Fqdn,
	serverPublicIpAddress valueObject.IpAddress,
) []valueObject.Fqdn {
	functionalHostnames := []valueObject.Fqdn{}

	for _, vhostName := range vhostNames {
		wwwVhostName, err := valueObject.NewFqdn("www." + vhostName.String())
		if err != nil {
			continue
		}

		vhostNames = append(vhostNames, wwwVhostName)
	}

	serverPublicIpAddressStr := serverPublicIpAddress.String()
	for _, vhostName := range vhostNames {
		vhostNameStr := vhostName.String()

		hostnameRecords, err := infraHelper.DnsLookup(vhostNameStr, nil)
		if err != nil || len(hostnameRecords) == 0 {
			continue
		}

		for _, record := range hostnameRecords {
			if record != serverPublicIpAddressStr {
				continue
			}

			functionalHostnames = append(functionalHostnames, vhostName)
		}
	}

	return functionalHostnames
}

func (repo *SslCmdRepo) createOwnershipValidationMapping(
	mappingCmdRepo *mappingInfra.MappingCmdRepo,
	targetVhostName valueObject.Fqdn,
	expectedOwnershipHash valueObject.Hash,
) (mappingId valueObject.MappingId, err error) {
	path, _ := valueObject.NewMappingPath(
		infraData.GlobalConfigs.DomainOwnershipValidationUrlPath,
	)
	matchPattern, _ := valueObject.NewMappingMatchPattern("equals")
	targetType, _ := valueObject.NewMappingTargetType("inline-html")
	httpResponseCode, _ := valueObject.NewHttpResponseCode(200)
	targetValue, _ := valueObject.NewMappingTargetValue(
		expectedOwnershipHash.String(), targetType,
	)

	inlineHtmlMapping := dto.NewCreateMapping(
		targetVhostName, path, matchPattern, targetType, &targetValue, &httpResponseCode,
	)

	mappingId, err = mappingCmdRepo.Create(inlineHtmlMapping)
	if err != nil {
		return mappingId, err
	}

	return mappingId, nil
}

func (repo *SslCmdRepo) httpFilterFunctionalHostnames(
	vhostNames []valueObject.Fqdn,
	expectedOwnershipHash valueObject.Hash,
	serverPublicIpAddress valueObject.IpAddress,
) []valueObject.Fqdn {
	functionalHostnames := []valueObject.Fqdn{}

	serverPublicIpAddressStr := serverPublicIpAddress.String()
	expectedHashStr := expectedOwnershipHash.String()
	mappingCmdRepo := mappingInfra.NewMappingCmdRepo(repo.persistentDbSvc)

	for _, vhostName := range vhostNames {
		vhostNameStr := vhostName.String()
		ownershipValidationMappingId, err := repo.createOwnershipValidationMapping(
			mappingCmdRepo, vhostName, expectedOwnershipHash,
		)
		if err != nil {
			continue
		}

		hashUrlPath := infraData.GlobalConfigs.DomainOwnershipValidationUrlPath
		hashUrlFull := "https://" + vhostNameStr + hashUrlPath
		curlBaseCmd := "curl -skLm 10 "
		sniFlag := "--resolve " + vhostNameStr + ":443:" + serverPublicIpAddressStr
		ownershipHashFound, err := infraHelper.RunCmdWithSubShell(
			curlBaseCmd + sniFlag + " " + hashUrlFull,
		)
		if err != nil {
			hashUrlFull = "https://" + serverPublicIpAddressStr + hashUrlPath
			ownershipHashFound, err = infraHelper.RunCmdWithSubShell(
				curlBaseCmd + "-H \"Host: " + vhostNameStr + "\" " + hashUrlFull,
			)
			if err != nil {
				continue
			}
		}

		if ownershipHashFound != expectedHashStr {
			continue
		}

		functionalHostnames = append(functionalHostnames, vhostName)

		err = mappingCmdRepo.Delete(ownershipValidationMappingId)
		if err != nil {
			log.Printf("DeleteOwnershipValidationMappingError: %s", err.Error())
		}
	}

	return functionalHostnames
}

func (repo *SslCmdRepo) issueValidSsl(
	mainHostname valueObject.Fqdn,
	functionalHostnames []valueObject.Fqdn,
) error {
	mainPairHostnameStr := mainHostname.String()
	vhostRootDir := infraData.GlobalConfigs.PrimaryPublicDir
	if !infraHelper.IsPrimaryVirtualHost(mainHostname) {
		vhostRootDir += "/" + mainPairHostnameStr
	}

	if !infraHelper.FileExists(vhostRootDir) {
		return errors.New("VirtualHostRootDirNotFound")
	}

	certbotCmd := "certbot certonly --webroot --webroot-path " + vhostRootDir +
		" --agree-tos --register-unsafely-without-email --cert-name " + mainPairHostnameStr
	for _, functionalHostname := range functionalHostnames {
		certbotCmd += " -d " + functionalHostname.String()
	}

	_, err := infraHelper.RunCmdWithSubShell(certbotCmd)
	if err != nil {
		return errors.New("GenerateValidSslCertError: " + err.Error())
	}

	certbotDirPath := "/etc/letsencrypt/live"
	shouldOverwrite := true

	certbotCrtFilePath := certbotDirPath + "/" + mainPairHostnameStr + "/fullchain.pem"
	vhostCrtFilePath := infraData.GlobalConfigs.PkiConfDir + "/" + mainPairHostnameStr + ".crt"
	err = infraHelper.CreateSymlink(certbotCrtFilePath, vhostCrtFilePath, shouldOverwrite)
	if err != nil {
		return errors.New("CreateSslCertSymlinkError: " + err.Error())
	}

	certbotKeyFilePath := certbotDirPath + "/" + mainPairHostnameStr + "/privkey.pem"
	vhostKeyFilePath := infraData.GlobalConfigs.PkiConfDir + "/" + mainPairHostnameStr + ".key"
	err = infraHelper.CreateSymlink(certbotKeyFilePath, vhostKeyFilePath, shouldOverwrite)
	if err != nil {
		return errors.New("CreateSslKeySymlinkError: " + err.Error())
	}

	return nil
}

func (repo *SslCmdRepo) ReplaceWithValidSsl(sslPair entity.SslPair) error {
	o11yQueryRepo := o11yInfra.NewO11yQueryRepo(repo.transientDbSvc)
	serverPublicIpAddress, err := o11yQueryRepo.ReadServerPublicIpAddress()
	if err != nil {
		return err
	}

	dnsFunctionalHostnames := repo.dnsFilterFunctionalHostnames(
		sslPair.VirtualHostsHostnames, serverPublicIpAddress,
	)
	if len(dnsFunctionalHostnames) == 0 {
		return errors.New("NoDnsFunctionalHostnames")
	}

	expectedOwnershipHash, err := repo.sslQueryRepo.GetOwnershipValidationHash(
		sslPair.Certificate.CertificateContent,
	)
	if err != nil {
		return errors.New(
			"CreateOwnershipValidationHashError: " + err.Error(),
		)
	}
	httpFunctionalHostnames := repo.httpFilterFunctionalHostnames(
		dnsFunctionalHostnames, expectedOwnershipHash, serverPublicIpAddress,
	)
	if len(httpFunctionalHostnames) == 0 {
		return errors.New("NoHttpFunctionalHostnames")
	}

	return repo.issueValidSsl(
		sslPair.VirtualHostsHostnames[0], httpFunctionalHostnames,
	)
}

func (repo *SslCmdRepo) Create(createSslPair dto.CreateSslPair) error {
	if len(createSslPair.VirtualHostsHostnames) == 0 {
		return errors.New("EmptyVirtualHosts")
	}

	firstVhostNameStr := createSslPair.VirtualHostsHostnames[0].String()
	firstVhostCertFilePath := infraData.GlobalConfigs.PkiConfDir + "/" + firstVhostNameStr + ".crt"
	firstVhostCertKeyFilePath := infraData.GlobalConfigs.PkiConfDir + "/" + firstVhostNameStr + ".key"

	for _, vhostName := range createSslPair.VirtualHostsHostnames {
		vhostStr := vhostName.String()
		vhostCertFilePath := infraData.GlobalConfigs.PkiConfDir + "/" + vhostStr + ".crt"
		vhostCertKeyFilePath := infraData.GlobalConfigs.PkiConfDir + "/" + vhostStr + ".key"

		shouldBeSymlink := vhostStr != firstVhostNameStr
		if shouldBeSymlink {
			shouldOverwrite := true
			err := infraHelper.CreateSymlink(
				firstVhostCertFilePath, vhostCertFilePath, shouldOverwrite,
			)
			if err != nil {
				log.Printf(
					"CreateSslCertSymlinkError (%s): %s", vhostName.String(), err.Error(),
				)
				continue
			}

			err = infraHelper.CreateSymlink(
				firstVhostCertKeyFilePath, vhostCertKeyFilePath, shouldOverwrite,
			)
			if err != nil {
				log.Printf(
					"CreateSslKeySymlinkError (%s): %s", vhostName.String(), err.Error(),
				)
				continue
			}

			continue
		}

		shouldOverwrite := true
		err := infraHelper.UpdateFile(
			vhostCertFilePath,
			createSslPair.Certificate.CertificateContent.String(),
			shouldOverwrite,
		)
		if err != nil {
			return err
		}

		err = infraHelper.UpdateFile(
			vhostCertKeyFilePath, createSslPair.Key.String(), shouldOverwrite,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *SslCmdRepo) Delete(sslId valueObject.SslId) error {
	sslPairToDelete, err := repo.sslQueryRepo.ReadById(sslId)
	if err != nil {
		return errors.New("SslNotFound")
	}

	for _, vhostName := range sslPairToDelete.VirtualHostsHostnames {
		err = repo.ReplaceWithSelfSigned(vhostName)
		if err != nil {
			log.Printf("%s (%s)", err.Error(), vhostName.String())
			continue
		}
	}

	return nil
}

func (repo *SslCmdRepo) DeleteSslPairVhosts(
	deleteDto dto.DeleteSslPairVhosts,
) error {
	vhostQueryRepo := vhostInfra.NewVirtualHostQueryRepo(repo.persistentDbSvc)
	for _, vhostName := range deleteDto.VirtualHostsHostnames {
		_, err := vhostQueryRepo.ReadByHostname(vhostName)
		if err != nil {
			continue
		}

		err = repo.ReplaceWithSelfSigned(vhostName)
		if err != nil {
			log.Printf(
				"DeleteSslPairVhostsError (%s): %s", vhostName.String(), err.Error(),
			)
		}
	}

	return nil
}
