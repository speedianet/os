package repository

import (
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/entity"
)

type VirtualHostCmdRepo interface {
	Create(createDto dto.CreateVirtualHost) error
	Delete(vhost entity.VirtualHost) error
	CreateMapping(createMapping dto.CreateMapping) error
	DeleteMapping(mapping entity.Mapping) error
	RecreateMapping(mapping entity.Mapping) error
}
