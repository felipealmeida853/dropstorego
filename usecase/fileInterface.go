package usecase

import "dropstore/dto"

type FileUseCaseInterface interface {
	GetFile(input dto.FileUseCaseInputDTO) (dto.FileUseCaseOutputDTO, error)
	SaveFile(input dto.FileUseCaseInputDTO) (dto.FileUseCaseOutputDTO, error)
	DeleteFile(input dto.FileUseCaseInputDTO) error
}