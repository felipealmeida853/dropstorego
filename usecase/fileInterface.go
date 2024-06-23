package usecase

import "dropstore/dto"

type FileUseCaseInterface interface {
	GetFile(input dto.FileUseCaseGetInputDTO) (dto.FileUseCaseGetOutputDTO, error)
	ListAll() ([]dto.FileRepositoryOutputDTO, error)
	SaveFile(input dto.FileUseCaseInputDTO) (dto.FileUseCaseOutputDTO, error)
	DeleteFile(input dto.FileUseCaseDeleteInputDTO) error
}
