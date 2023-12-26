package useCase

import (
	"errors"
	"log"

	"github.com/speedianet/os/src/domain/dto"
	"github.com/speedianet/os/src/domain/repository"
)

func CompressUnixFiles(
	filesQueryRepo repository.FilesQueryRepo,
	filesCmdRepo repository.FilesCmdRepo,
	compressUnixFiles dto.CompressUnixFiles,
) (dto.CompressionProcessReport, error) {
	compressionProcessReport := filesCmdRepo.Compress(compressUnixFiles)

	allPathsFailedCompression := len(compressionProcessReport.FilePathsThatFailedToCompressWithReason) == len(compressUnixFiles.SourcePaths)
	if allPathsFailedCompression {
		return compressionProcessReport, errors.New("CompressUnixFilesInfraError")
	}

	log.Printf("Compressed file %s created.", compressUnixFiles.DestinationPath)

	return compressionProcessReport, nil
}
