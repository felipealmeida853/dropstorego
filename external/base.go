package external

import "dropstore/dto"

type FileStoreBucketInterface interface {
	Save(input dto.FileStoreBucketInputDTO) (dto.FileStoreBucketOutputDTO, error)
	Get(input dto.FileStoreBucketInputDTO) (dto.FileStoreBucketOutputDTO, error)
	Delete() error
}
