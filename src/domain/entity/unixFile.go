package entity

import (
	"github.com/speedianet/os/src/domain/valueObject"
)

type UnixFile struct {
	Name        valueObject.UnixFileName        `json:"name"`
	Path        valueObject.UnixFilePath        `json:"path"`
	MimeType    valueObject.MimeType            `json:"mimeType"`
	Permissions valueObject.UnixFilePermissions `json:"permissions"`
	Size        valueObject.Byte                `json:"size"`
	Extension   *valueObject.UnixFileExtension  `json:"extension"`
	Uid         valueObject.UnixUid             `json:"uid"`
	Owner       valueObject.Username            `json:"owner"`
	Gid         valueObject.GroupId             `json:"gid"`
	Group       valueObject.GroupName           `json:"group"`
	UpdatedAt   valueObject.UnixTime            `json:"updatedAt"`
}

func NewUnixFile(
	name valueObject.UnixFileName,
	path valueObject.UnixFilePath,
	mimeType valueObject.MimeType,
	permissions valueObject.UnixFilePermissions,
	size valueObject.Byte,
	extension *valueObject.UnixFileExtension,
	uid valueObject.UnixUid,
	owner valueObject.Username,
	gid valueObject.GroupId,
	group valueObject.GroupName,
	updatedAt valueObject.UnixTime,
) UnixFile {
	return UnixFile{
		Name:        name,
		Path:        path,
		MimeType:    mimeType,
		Permissions: permissions,
		Size:        size,
		Extension:   extension,
		Uid:         uid,
		Owner:       owner,
		Gid:         gid,
		Group:       group,
		UpdatedAt:   updatedAt,
	}
}
