package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/repository"
	"github.com/speedianet/os/src/domain/valueObject"
)

func UploadUnixFiles(
	filesQueryRepo repository.FilesQueryRepo,
	filesCmdRepo repository.FilesCmdRepo,
	uploadUnixFiles dto.UploadUnixFiles,
) (dto.UploadProcessReport, error) {
	filesLargerThanAllowedFailure := []valueObject.UploadProcessFailure{}
	filesWithAllowedSizes := []valueObject.FileStreamHandler{}
	largerFileErrMessage := "File size is greater than 5 GB"
	for _, fileToUploadStream := range uploadUnixFiles.FileStreamHandlers {
		fileStreamHandlerSizeInGB := fileToUploadStream.Size.ToGiB()
		if fileStreamHandlerSizeInGB > 5 {
			log.Printf("UploadUnixFileError: %s", largerFileErrMessage)

			failureReason, _ := valueObject.NewProcessFileFailure(largerFileErrMessage)

			filesLargerThanAllowedFailure = append(
				filesLargerThanAllowedFailure,
				valueObject.NewUploadProcessFailure(
					fileToUploadStream.Name,
					failureReason,
				),
			)

			continue
		}

		filesWithAllowedSizes = append(filesWithAllowedSizes, fileToUploadStream)
	}

	uploadUnixFiles.FileStreamHandlers = filesWithAllowedSizes

	uploadProcessReport := filesCmdRepo.Upload(uploadUnixFiles)

	uploadProcessReport.FailedPathsWithReason = append(
		uploadProcessReport.FailedPathsWithReason,
		filesLargerThanAllowedFailure...,
	)

	failedPathsWithReasonLen := len(uploadProcessReport.FailedPathsWithReason)
	fileStreamHandlersLen := len(uploadUnixFiles.FileStreamHandlers)
	allFilesFailedToUpload := failedPathsWithReasonLen == fileStreamHandlersLen
	if allFilesFailedToUpload {
		return uploadProcessReport, errors.New("UploadUnixFileInfraError")
	}

	log.Printf("Files uploaded to '%s'.", uploadUnixFiles.DestinationPath)

	return uploadProcessReport, nil
}
