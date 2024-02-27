package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/repository"
)

func CreateAccount(
	accQueryRepo repository.AccQueryRepo,
	accCmdRepo repository.AccCmdRepo,
	filesQueryRepo repository.FilesQueryRepo,
	filesCmdRepo repository.FilesCmdRepo,
	createAccount dto.CreateAccount,
) error {
	_, err := accQueryRepo.GetByUsername(createAccount.Username)
	if err == nil {
		return errors.New("UsernameAlreadyExists")
	}

	err = accCmdRepo.Create(createAccount)
	if err != nil {
		return errors.New("CreateAccountError")
	}

	log.Printf("Account '%v' created.", createAccount.Username.String())

	return CreateTrash(filesQueryRepo, filesCmdRepo)
}
