package repository

import "dropstore/dto"

type FileRepositoryInterface interface {
	Save(input dto.FileRepositoryInputDTO) (dto.FileRepositoryOutputDTO, error)
	GetByKey(input dto.FileRepositoryKeyInputDTO) (dto.FileRepositoryOutputDTO, error)
	DeleteByKey(input dto.FileRepositoryKeyInputDTO) error
	ListAll() ([]dto.FileRepositoryOutputDTO, error)
}
