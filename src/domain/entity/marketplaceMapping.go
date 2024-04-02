package entity

import (
	"github.com/speedianet/os/src/domain/valueObject"
	"gopkg.in/yaml.v3"
)

type MarketplaceMapping struct {
	Path              valueObject.MappingPath         `json:"path"`
	MatchPattern      valueObject.MappingMatchPattern `json:"matchPattern"`
	TargetServiceName *valueObject.ServiceName        `json:"targetServiceName"`
}

func NewMarketplaceMapping(
	path valueObject.MappingPath,
	matchPattern valueObject.MappingMatchPattern,
	targetServiceName *valueObject.ServiceName,
) MarketplaceMapping {
	return MarketplaceMapping{
		Path:              path,
		MatchPattern:      matchPattern,
		TargetServiceName: targetServiceName,
	}
}

func (mmPtr *MarketplaceMapping) UnmarshalYAML(value *yaml.Node) error {
	var valuesMap map[string]string
	err := value.Decode(&valuesMap)
	if err != nil {
		return err
	}

	path, err := valueObject.NewMappingPath(valuesMap["path"])
	if err != nil {
		return err
	}

	matchPattern, err := valueObject.NewMappingMatchPattern(
		valuesMap["matchPattern"],
	)
	if err != nil {
		return err
	}

	var targetSvcNamePtr *valueObject.ServiceName
	targetSvcName, targetSvcNameExists := valuesMap["targetServiceName"]
	if targetSvcNameExists {
		targetSvcName, err := valueObject.NewServiceName(targetSvcName)
		if err != nil {
			return err
		}

		targetSvcNamePtr = &targetSvcName
	}

	*mmPtr = NewMarketplaceMapping(path, matchPattern, targetSvcNamePtr)

	return nil
}