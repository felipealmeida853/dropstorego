package external

import (
	"context"
	"dropstore/dto"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type FileStoreBucketS3 struct {
	ctx        context.Context
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	svc        *s3.S3
}

func NewFileStoreBucketS3(ctx context.Context, uploader *s3manager.Uploader, downloader *s3manager.Downloader, svc *s3.S3) *FileStoreBucketS3 {
	return &FileStoreBucketS3{ctx, uploader, downloader, svc}
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
		Bucket: aws.String(input.Bucket),
		Key:    aws.String(input.Key),
		Body:   file,
	})
	if err != nil {
		fmt.Printf("Error uploading file to bucket: %s, filename: %s", input.Bucket, input.Filename)
		return result, err
	}
	result.Filename = input.Filename
	result.Key = input.Key
	return result, nil
}

func (s *FileStoreBucketS3) Get(input dto.FileStoreBucketInputDTO) (dto.FileStoreBucketOutputDTO, error) {
	var result dto.FileStoreBucketOutputDTO
	filePath := "./" + input.Filename
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file, err: %v", err)
		return result, err
	}

	numBytes, err := s.downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(input.Bucket),
			Key:    aws.String(input.Key),
		})
	if err != nil {
		fmt.Printf("Error downloading on bucket, err: %v", err)
		return result, err
	}
	fmt.Printf("Number of bytes downloaded %d", numBytes)
	result.Filename = file.Name()
	result.File = file
	result.FilePath = filePath
	return result, nil
}

func (s *FileStoreBucketS3) Delete(input dto.FileStoreBucketDeleteInputDTO) error {
	_, err := s.svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(input.Bucket),
		Key:    aws.String(input.Key),
	})
	if err != nil {
		fmt.Printf("Error deleting file key: %s,  err: %v", input.Key, err)
		return err
	}
	return nil
}
