package external

import (
	"context"
	"dropstore/dto"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileStoreBucketS3 struct {
	ctx      context.Context
	uploader *s3manager.Uploader
}

func NewFileStoreBucketS3(ctx context.Context, uploader *s3manager.Uploader) *FileStoreBucketS3 {
	return &FileStoreBucketS3{ctx, uploader}
}

func (s *FileStoreBucketS3) Save(input dto.FileStoreBucketInputDTO) (dto.FileStoreBucketOutputDTO, error) {
	var result dto.FileStoreBucketOutputDTO
	file, err := os.Open(input.Path)
	if err != nil {
		fmt.Printf("Error opening file: %s, error: %v", input.Filename, err)
		return result, err
	}
	defer file.Close()

	_, err = s.uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(input.BucketName),
		Key:    aws.String(input.Key),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Error uploading file to bucket: %s, filename: %s", input.BucketName, input.Filename)
		return result, err
	}
	result.Filename = input.Filename
	result.Key = input.Key
	return result, nil
}

func (s *FileStoreBucketS3) Get(input dto.FileStoreBucketInputDTO) (dto.FileStoreBucketOutputDTO, error) {
	//TODO: Implement external s3
	var result dto.FileStoreBucketOutputDTO
	return result, nil
}

func (s *FileStoreBucketS3) Delete() error {
	//TODO: Implement external s3
	return nil
}
