package repository

import (
	"github.com/speedianet/os/src/domain/entity"
	"github.com/speedianet/os/src/domain/valueObject"
)

type FilesQueryRepo interface {
	Exists(unixFilePath valueObject.UnixFilePath) bool
	IsDir(unixFilePath valueObject.UnixFilePath) (bool, error)
	Get(unixFilePath valueObject.UnixFilePath) ([]entity.UnixFile, error)
}
