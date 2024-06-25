package repository

import "dropstore/dto"

type FileRepositoryInterface interface {
	Save(input dto.FileRepositoryInputDTO) (dto.FileRepositoryOutputDTO, error)
	GetByKey(input dto.FileRepositoryKeyInputDTO) (dto.FileRepositoryOutputDTO, error)
	DeleteByKey(input dto.FileRepositoryKeyInputDTO) error
	ListAll(dto.FileRepositoryListAllInputDTO) ([]dto.FileRepositoryOutputDTO, error)
}

type FolderRepositoryInterface interface {
	Save(input dto.FolderRepositoryInputDTO) (dto.FolderRepositoryOutputDTO, error)
	ListAll() ([]dto.FolderRepositoryOutputDTO, error)
}
