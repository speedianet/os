package repository

import (
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/valueObject"
)

type FilesCmdRepo interface {
	Create(dto.AddUnixFile) error
	Move(dto.UpdateUnixFile) error
	Copy(dto.CopyUnixFile) error
	UpdateContent(dto.UpdateUnixFileContent) error
	UpdatePermissions(
		valueObject.UnixFilePath,
		valueObject.UnixFilePermissions,
	) error
	Delete([]valueObject.UnixFilePath)
	Compress(dto.CompressUnixFiles) dto.CompressionProcessReport
	Extract(dto.ExtractUnixFiles) error
	Upload(uploadUnixFiles dto.UploadUnixFiles) dto.UploadProcessReport
}
