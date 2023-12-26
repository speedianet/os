package infra

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/valueObject"
	infraHelper "github.com/speedianet/os/src/infra/helper"
)

type FilesCmdRepo struct{}

func fillUploadProcessReportFailure(
	currentUploadProcessReportList []valueObject.UploadProcessFailure,
	errMessage string,
	fileStreamHandlers []valueObject.FileStreamHandler,
) []valueObject.UploadProcessFailure {
	uploadProcessReportList := currentUploadProcessReportList

	uploadProcessReportList = append(
		uploadProcessReportList,
		uploadProcessReportFailureListFactory(errMessage, fileStreamHandlers)...,
	)

	return uploadProcessReportList
}

func uploadProcessReportFailureListFactory(
	errMessage string,
	fileStreamHandlers []valueObject.FileStreamHandler,
) []valueObject.UploadProcessFailure {
	uploadProcessReportFailureList := []valueObject.UploadProcessFailure{}

	for _, fileStreamHandler := range fileStreamHandlers {
		uploadProcessReportFailureList = append(
			uploadProcessReportFailureList,
			valueObject.NewUploadProcessFailure(fileStreamHandler.GetFileName(), errMessage),
		)
	}

	return uploadProcessReportFailureList
}

func (repo FilesCmdRepo) Create(addUnixFile dto.AddUnixFile) error {
	filesExists := infraHelper.FileExists(addUnixFile.Path.String())
	if filesExists {
		return errors.New("PathAlreadyExists")
	}

	if !addUnixFile.Type.IsDir() {
		_, err := os.Create(addUnixFile.Path.String())
		if err != nil {
			log.Printf("CreateUnixFileError: %s", err)
			return errors.New("CreateUnixFileError")
		}

		return repo.UpdatePermissions(
			addUnixFile.Path,
			addUnixFile.Permissions,
		)
	}

	err := os.MkdirAll(addUnixFile.Path.String(), addUnixFile.Permissions.GetFileMode())
	if err != nil {
		log.Printf("CreateUnixDirectoryError: %s", err)
		return errors.New("CreateUnixDirectoryError")
	}

	return nil
}

func (repo FilesCmdRepo) Move(updateUnixFile dto.UpdateUnixFile) error {
	fileToMoveExists := infraHelper.FileExists(updateUnixFile.Path.String())
	if !fileToMoveExists {
		return errors.New("FileToMoveDoesNotExists")
	}

	destinationPathExists := infraHelper.FileExists(updateUnixFile.DestinationPath.String())
	if destinationPathExists {
		return errors.New("DestinationPathAlreadyExists")
	}

	return os.Rename(
		updateUnixFile.Path.String(),
		updateUnixFile.DestinationPath.String(),
	)
}

func (repo FilesCmdRepo) Copy(copyUnixFile dto.CopyUnixFile) error {
	fileToCopyExists := infraHelper.FileExists(copyUnixFile.OriginPath.String())
	if !fileToCopyExists {
		return errors.New("FileToCopyDoesNotExists")
	}

	destinationPathExists := infraHelper.FileExists(copyUnixFile.DestinationPath.String())
	if destinationPathExists {
		return errors.New("DestinationPathAlreadyExists")
	}

	_, err := infraHelper.RunCmd(
		"rsync",
		"-avq",
		copyUnixFile.OriginPath.String(),
		copyUnixFile.DestinationPath.String(),
	)
	return err
}

func (repo FilesCmdRepo) UpdateContent(
	updateUnixFileContent dto.UpdateUnixFileContent,
) error {
	queryRepo := FilesQueryRepo{}

	fileToUpdateContent, err := queryRepo.GetOnly(updateUnixFileContent.Path)
	if err != nil {
		return err
	}

	if fileToUpdateContent.MimeType.IsDir() {
		return errors.New("PathIsADir")
	}

	return infraHelper.UpdateFile(
		updateUnixFileContent.Path.String(),
		updateUnixFileContent.Content.GetDecodedContent(),
		true,
	)
}

func (repo FilesCmdRepo) UpdatePermissions(
	unixFilePath valueObject.UnixFilePath,
	unixFilePermissions valueObject.UnixFilePermissions,
) error {
	queryRepo := FilesQueryRepo{}

	_, err := queryRepo.Get(unixFilePath)
	if err != nil {
		return err
	}

	return os.Chmod(unixFilePath.String(), unixFilePermissions.GetFileMode())
}

func (repo FilesCmdRepo) Compress(
	compressUnixFiles dto.CompressUnixFiles,
) dto.CompressionProcessReport {
	compressBinary := "tar"
	compressBinaryFlag := "-czf"
	compressExtFile := ".tar.gz"
	if compressUnixFiles.CompressionType.String() == "zip" {
		compressBinary = "zip"
		compressBinaryFlag = "-qr"
		compressExtFile = ".zip"
	}

	destinationPathStr := compressUnixFiles.DestinationPath.String()

	destinationPathExtension := compressUnixFiles.DestinationPath.GetFileExtension()
	if !destinationPathExtension.IsEmpty() {
		destinationPathWithoutExt := strings.Split(destinationPathStr, ".")[0]
		destinationPathStr = destinationPathWithoutExt
	}

	destinationPathWithCompressionTypeAsExtStr := destinationPathStr + compressExtFile
	destinationPathWithCompressionTypeAsExt, _ := valueObject.NewUnixFilePath(destinationPathWithCompressionTypeAsExtStr)

	compressionProcessReport := dto.NewCompressionProcessReport(
		[]valueObject.UnixFilePath{},
		[]valueObject.CompressionProcessFailure{},
		destinationPathWithCompressionTypeAsExt,
	)

	destinationPathExists := infraHelper.FileExists(destinationPathWithCompressionTypeAsExtStr)
	if destinationPathExists {
		errMessage := "DestinationPathAlreadyExists"
		for _, failedFile := range compressUnixFiles.Paths {
			compressionProcessReport.Failure = append(
				compressionProcessReport.Failure,
				valueObject.NewCompressionProcessFailure(failedFile, errMessage),
			)
		}

		return compressionProcessReport
	}

	filesToCompressStrList := []string{}

	for _, fileToCompress := range compressUnixFiles.Paths {
		fileToCompressExists := infraHelper.FileExists(fileToCompress.String())
		if !fileToCompressExists {
			compressionProcessReport.Failure = append(
				compressionProcessReport.Failure,
				valueObject.NewCompressionProcessFailure(fileToCompress, "FileDoesNotExists"),
			)

			continue
		}

		compressionProcessReport.Success = append(compressionProcessReport.Success, fileToCompress)
		filesToCompressStrList = append(filesToCompressStrList, fileToCompress.String())
	}

	if len(compressionProcessReport.Success) < 1 {
		return compressionProcessReport
	}

	filesToCompressStr := strings.Join(filesToCompressStrList, " ")

	_, err := infraHelper.RunCmd(
		compressBinary,
		compressBinaryFlag,
		destinationPathWithCompressionTypeAsExtStr,
		filesToCompressStr,
	)

	if err != nil {
		for _, fileThatFailedCompression := range compressionProcessReport.Success {
			compressionProcessReport.Failure = append(
				compressionProcessReport.Failure,
				valueObject.NewCompressionProcessFailure(fileThatFailedCompression, err.Error()),
			)
		}

		compressionProcessReport.Success = []valueObject.UnixFilePath{}
	}

	return compressionProcessReport
}

func (repo FilesCmdRepo) Extract(extractUnixFiles dto.ExtractUnixFiles) error {
	fileToExtract := extractUnixFiles.Path

	fileToExtractExists := infraHelper.FileExists(fileToExtract.String())
	if !fileToExtractExists {
		return errors.New("FileToExtractDoesNotExists")
	}

	destinationPath := extractUnixFiles.DestinationPath

	destinationPathExists := infraHelper.FileExists(destinationPath.String())
	if destinationPathExists {
		return errors.New("DEstinationPathAlreadyExists")
	}

	compressBinary := "tar"
	compressBinaryFlag := "-xf"
	compressDestinationFlag := "-C"

	unixFilePathExtension := fileToExtract.GetFileExtension()
	if unixFilePathExtension.String() == "zip" {
		compressBinary = "unzip"
		compressBinaryFlag = "-qq"
		compressDestinationFlag = "-d"
	}

	err := infraHelper.MakeDir(destinationPath.String())
	if err != nil {
		return err
	}

	_, err = infraHelper.RunCmd(
		compressBinary,
		compressBinaryFlag,
		fileToExtract.String(),
		compressDestinationFlag,
		destinationPath.String(),
	)

	return err
}

func (repo FilesCmdRepo) Delete(
	unixFilePath []valueObject.UnixFilePath,
) {
	for _, fileToDelete := range unixFilePath {
		fileExists := infraHelper.FileExists(fileToDelete.String())
		if !fileExists {
			log.Printf("DeleteFileError: FileDoesNotExists")
			continue
		}

		err := os.RemoveAll(fileToDelete.String())
		if err != nil {
			log.Printf("DeleteFileError: %s", err)
			continue
		}

		log.Printf("File '%s' deleted.", fileToDelete.String())
	}
}

func (repo FilesCmdRepo) Upload(
	uploadUnixFiles dto.UploadUnixFiles,
) dto.UploadProcessReport {
	queryRepo := FilesQueryRepo{}

	destinationPath := uploadUnixFiles.DestinationPath

	uploadProcessReport := dto.NewUploadProcessReport(
		[]valueObject.UnixFileName{},
		[]valueObject.UploadProcessFailure{},
		destinationPath,
	)

	destinationFile, err := queryRepo.GetOnly(destinationPath)
	if err != nil {
		uploadProcessReport.Failure = fillUploadProcessReportFailure(
			uploadProcessReport.Failure,
			err.Error(),
			uploadUnixFiles.FileStreamHandlers,
		)

		return uploadProcessReport
	}

	if !destinationFile.MimeType.IsDir() {
		uploadProcessReport.Failure = fillUploadProcessReportFailure(
			uploadProcessReport.Failure,
			"DestinationPathCannotBeAFile",
			uploadUnixFiles.FileStreamHandlers,
		)

		return uploadProcessReport
	}

	for _, fileToUpload := range uploadUnixFiles.FileStreamHandlers {
		destinationFilePath := destinationPath.String() + "/" + fileToUpload.GetFileName().String()
		destinationEmptyFile, err := os.Create(destinationFilePath)
		if err != nil {
			errMessage := fmt.Sprintf("CreateEmptyFileToStoreUploadFileError: %s", err.Error())
			uploadProcessReport.Failure = fillUploadProcessReportFailure(
				uploadProcessReport.Failure,
				errMessage,
				uploadUnixFiles.FileStreamHandlers,
			)

			continue
		}
		defer destinationEmptyFile.Close()

		fileToUploadStream, err := fileToUpload.Open()
		if err != nil {
			errMessage := fmt.Sprintf("UnableToOpenFileStream: %s", err.Error())
			uploadProcessReport.Failure = fillUploadProcessReportFailure(
				uploadProcessReport.Failure,
				errMessage,
				uploadUnixFiles.FileStreamHandlers,
			)

			continue
		}

		_, err = io.Copy(destinationEmptyFile, fileToUploadStream)
		if err != nil {
			errMessage := fmt.Sprintf("CopyFileStreamHandlerContentToDestinationFileError: %s", err.Error())
			uploadProcessReport.Failure = fillUploadProcessReportFailure(
				uploadProcessReport.Failure,
				errMessage,
				uploadUnixFiles.FileStreamHandlers,
			)

			continue
		}

		uploadProcessReport.Success = append(uploadProcessReport.Success, fileToUpload.GetFileName())
	}

	return uploadProcessReport
}
