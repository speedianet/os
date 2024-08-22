package valueObject

import (
	"errors"
	"slices"
	"strings"

	voHelper "github.com/speedianet/os/src/domain/valueObject/helper"
)

type MappingTargetType string

var validMappingTargetTypes = []string{
	"url", "service", "response-code", "inline-html", "static-files",
}

func NewMappingTargetType(value interface{}) (
	mappingTargetType MappingTargetType, err error,
) {
	stringValue, err := voHelper.InterfaceToString(value)
	if err != nil {
		return mappingTargetType, errors.New("MappingTargetTypeMustBeString")
	}
	stringValue = strings.ToLower(stringValue)

	if !slices.Contains(validMappingTargetTypes, stringValue) {
		return mappingTargetType, errors.New("InvalidMappingTargetType")
	}

	return MappingTargetType(stringValue), nil
}

func (vo MappingTargetType) String() string {
	return string(vo)
}
