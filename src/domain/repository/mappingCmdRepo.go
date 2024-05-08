package repository

import (
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/entity"
	"github.com/speedianet/os/src/domain/valueObject"
)

type MappingCmdRepo interface {
	Create(createDto dto.CreateMapping) (valueObject.MappingId, error)
	Delete(mappingId valueObject.MappingId) error
	DeleteAuto(serviceName valueObject.ServiceName) error
	Recreate(mapping entity.Mapping) error
}