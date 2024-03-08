package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/repository"
	"github.com/speedianet/os/src/domain/valueObject"
)

func CreateMapping(
	queryRepo repository.VirtualHostQueryRepo,
	cmdRepo repository.VirtualHostCmdRepo,
	svcsQueryRepo repository.ServicesQueryRepo,
	createMapping dto.CreateMapping,
) error {
	vhostWithMappings, err := queryRepo.GetWithMappings()
	if err != nil {
		log.Printf("GetVirtualHostsError: %s", err.Error())
		return errors.New("GetVirtualHostsInfraError")
	}

	vhostIndex := -1
	for vhostWithMappingIndex, vhostWithMapping := range vhostWithMappings {
		if vhostWithMapping.Hostname != createMapping.Hostname {
			continue
		}

		for _, mapping := range vhostWithMapping.Mappings {
			if mapping.MatchPattern != createMapping.MatchPattern {
				continue
			}

			if mapping.Path != createMapping.Path {
				continue
			}

			return errors.New("MappingAlreadyExists")
		}

		vhostIndex = vhostWithMappingIndex
	}

	if vhostIndex == -1 {
		return errors.New("VirtualHostNotFound")
	}

	if vhostWithMappings[vhostIndex].Type.String() == "alias" {
		return errors.New("AliasCannotHaveMappings")
	}

	isServiceTarget := createMapping.TargetType.String() == "service"
	if isServiceTarget {
		if createMapping.TargetServiceName == nil {
			return errors.New("TargetServiceNameRequired")
		}

		service, err := svcsQueryRepo.GetByName(*createMapping.TargetServiceName)
		if err != nil {
			return err
		}

		if len(service.PortBindings) == 0 {
			return errors.New("ServiceDoesNotExposeAnyPorts")
		}

		for _, portBinding := range service.PortBindings {
			isTcp := portBinding.Protocol.String() == "tcp"
			isUdp := portBinding.Protocol.String() == "udp"
			if isTcp || isUdp {
				return errors.New("Layer4MappingNotSupportedYet")
			}
		}
	}

	isUrlTarget := createMapping.TargetType.String() == "url"
	if isUrlTarget && createMapping.TargetUrl == nil {
		return errors.New("TargetUrlRequired")
	}

	defaultUrlResponseCode, _ := valueObject.NewHttpResponseCode(301)
	if isUrlTarget && createMapping.TargetHttpResponseCode == nil {
		createMapping.TargetHttpResponseCode = &defaultUrlResponseCode
	}

	isStaticFiles := createMapping.TargetType.String() == "static-files"
	defaultStaticFilesResponseCode, _ := valueObject.NewHttpResponseCode(404)
	if isStaticFiles && createMapping.TargetHttpResponseCode == nil {
		createMapping.TargetHttpResponseCode = &defaultStaticFilesResponseCode
	}

	isResponseCodeTarget := createMapping.TargetType.String() == "response-code"
	if isResponseCodeTarget && createMapping.TargetHttpResponseCode == nil {
		return errors.New("TargetHttpResponseCodeRequired")
	}

	err = cmdRepo.CreateMapping(createMapping)
	if err != nil {
		log.Printf("CreateMappingError: %s", err.Error())
		return errors.New("CreateMappingInfraError")
	}

	return nil
}
