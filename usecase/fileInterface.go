package usecase

import "dropstore/dto"

type FileUseCaseInterface interface {
	GetFile(input dto.FileUseCaseGetInputDTO) (dto.FileUseCaseGetOutputDTO, error)
	ListAll(input dto.FileUseCaseListAllInputDTO) ([]dto.FileRepositoryOutputDTO, error)
	SaveFile(input dto.FileUseCaseInputDTO) (dto.FileUseCaseOutputDTO, error)
	DeleteFile(input dto.FileUseCaseDeleteInputDTO) error
}
