package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/repository"
)

func AddUnixFile(
	filesQueryRepo repository.FilesQueryRepo,
	filesCmdRepo repository.FilesCmdRepo,
	addUnixFile dto.AddUnixFile,
) error {
	fileType := addUnixFile.Type.GetWithFirstLetterUpperCase()

	unixFileExists, err := filesQueryRepo.Exists(addUnixFile.Path)
	if err != nil {
		return err
	}

	if unixFileExists {
		return errors.New(fileType + "AlreadyExists")
	}

	err = filesCmdRepo.Create(addUnixFile)
	if err != nil {
		return errors.New("Create" + fileType + "Error")
	}

	log.Printf(
		"%s '%s' created to '%s'.",
		fileType,
		addUnixFile.Path.GetFileName().String(),
		addUnixFile.Path.GetFileDir().String(),
	)

	return nil
}
