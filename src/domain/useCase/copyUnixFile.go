package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/repository"
)

func CopyUnixFile(
	filesQueryRepo repository.FilesQueryRepo,
	filesCmdRepo repository.FilesCmdRepo,
	copyUnixFile dto.CopyUnixFile,
) error {
	filePath := copyUnixFile.OriginPath

	fileIsDir, err := filesQueryRepo.IsDir(filePath)
	if err != nil {
		return err
	}

	fileExists, err := filesQueryRepo.Exists(filePath)
	if err != nil {
		return err
	}

	if fileExists {
		return errors.New("PathAlreadyExists")
	}

	fileType := "File"
	if fileIsDir {
		fileType = "Dir"
	}

	err = filesCmdRepo.Copy(copyUnixFile)
	if err != nil {
		log.Printf("Add%sCopyError: %s", fileType, err.Error())
		return errors.New("Add" + fileType + "CopyError")
	}

	fileOriginPath := copyUnixFile.OriginPath
	fileDestinationPath := copyUnixFile.DestinationPath
	log.Printf(
		"%s '%s' (%s) copy added to '%s' with name '%s'.",
		fileType,
		fileOriginPath.GetFileName().String(),
		fileOriginPath.GetFileDir().String(),
		fileDestinationPath.GetFileName().String(),
		fileDestinationPath.GetFileDir().String(),
	)

	return nil
}
