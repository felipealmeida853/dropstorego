package usecase

import "dropstore/dto"

type FolderUseCaseInterface interface {
	CreateFolder(input dto.FolderUseCaseCreateInputDTO) (dto.FolderUseCaseCreateOutputDTO, error)
	ListAll() ([]dto.FolderRepositoryOutputDTO, error)
}
