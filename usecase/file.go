package usecase

import (
	"dropstore/config"
	"dropstore/dto"
	"dropstore/external"
	"dropstore/repository"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type FileUseCase struct {
	repository      repository.FileRepositoryInterface
	fileStoreBucket external.FileStoreBucketInterface
	config          *config.Config
}

func NewFileUseCase(repository repository.FileRepositoryInterface, fileStoreBucket external.FileStoreBucketInterface, config *config.Config) *FileUseCase {
	return &FileUseCase{repository, fileStoreBucket, config}
}

func (uc *FileUseCase) GetFile(input dto.FileUseCaseGetInputDTO) (dto.FileUseCaseGetOutputDTO, error) {
	var result dto.FileUseCaseGetOutputDTO

	fileDB, err := uc.repository.GetByKey(dto.FileRepositoryKeyInputDTO{
		Key: input.Key,
	})

	fileStoreDTO, err := uc.fileStoreBucket.Get(dto.FileStoreBucketInputDTO{
		Filename: fileDB.Name,
		Key:      input.Key,
		Bucket:   fileDB.Bucket,
	})
	if err != nil {
		fmt.Printf("Error getting file, err: %v", err)
		return result, err
	}

	result.File = fileStoreDTO.File
	result.User = fileDB.User
	result.Name = fileDB.Name
	result.Bucket = fileDB.Bucket
	return result, nil
}

func (uc *FileUseCase) SaveFile(input dto.FileUseCaseInputDTO) (dto.FileUseCaseOutputDTO, error) {
	var result dto.FileUseCaseOutputDTO
	key := uuid.New().String() + "_" + input.Filename

	fileInfo, err := os.Stat(input.Path)
	if err != nil {
		fmt.Printf("Error on get stats of file, error: %v", err)
	}

	resultSaveFileOnBucket, err := uc.fileStoreBucket.Save(dto.FileStoreBucketInputDTO{
		Key:      key,
		Path:     input.Path,
		Filename: input.Filename,
		Bucket:   uc.config.BucketName,
	})
	if err != nil {
		fmt.Printf("UC Error saving on bucket, filename: %v", input.Filename)
		return result, err
	}

	err = os.Remove(resultSaveFileOnBucket.Filename)
	if err != nil {
		fmt.Printf("Error removing filename: %v", resultSaveFileOnBucket.Filename)
	}

	insertDB, err := uc.repository.Save(dto.FileRepositoryInputDTO{
		Name:      input.Filename,
		User:      input.User,
		Key:       key,
		CreatedAt: time.Now(),
		Bucket:    uc.config.BucketName,
		SizeInMB:  fileInfo.Size() / (1 << 20),
		Size:      fileInfo.Size(),
	})
	if err != nil {
		fmt.Printf("UC Error saving on DB, filename: %v", input.Filename)
		return result, err
	}

	result.Name = insertDB.Name
	result.User = insertDB.User
	result.Bucket = insertDB.Bucket
	result.CreatedAt = insertDB.CreatedAt
	result.Key = insertDB.Key
	result.SizeInMB = insertDB.SizeInMB
	result.SizeInMB = insertDB.Size

	return result, nil
}

func (uc *FileUseCase) DeleteFile(input dto.FileUseCaseInputDTO) error {
	//TODO: Implement uc
	return nil
}
