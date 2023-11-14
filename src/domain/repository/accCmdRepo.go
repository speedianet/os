package repository

import (
	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/valueObject"
)

type AccCmdRepo interface {
	Add(addUser dto.AddAccount) error
	Delete(accId valueObject.AccountId) error
	UpdatePassword(accId valueObject.AccountId, password valueObject.Password) error
	UpdateApiKey(accId valueObject.AccountId) (valueObject.AccessTokenStr, error)
}
