package usecase

import (
	"dropstore/dto"
	"dropstore/repository"
	"fmt"

	"github.com/google/uuid"
)

type FolderUseCase struct {
	repository repository.FolderRepositoryInterface
}

func NewFolderUseCase(repository repository.FolderRepositoryInterface) *FolderUseCase {
	return &FolderUseCase{repository}
}

func (uc *FolderUseCase) CreateFolder(input dto.FolderUseCaseCreateInputDTO) (dto.FolderUseCaseCreateOutputDTO, error) {
	var result dto.FolderUseCaseCreateOutputDTO

	insertDB, err := uc.repository.Save(dto.FolderRepositoryInputDTO{
		UUID: uuid.NewString(),
		Name: input.Name,
		User: input.User,
	})
	if err != nil {
		fmt.Printf("UC Error saving folder on DB, name: %v", input.Name)
		return result, err
	}

	result.Name = insertDB.Name
	result.UUID = insertDB.UUID

	return result, nil
}

func (uc *FolderUseCase) ListAll() ([]dto.FolderRepositoryOutputDTO, error) {
	result, err := uc.repository.ListAll()
	if err != nil {
		fmt.Printf("Error Listing all folders from DB err: %v", err)
		return nil, err
	}
	return result, err
}
