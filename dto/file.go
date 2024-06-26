package dto

import (
	"os"
	"time"
)

type FileUseCaseDeleteInputDTO struct {
	Key string `json:"key" bson:"key"`
}

type FileUseCaseInputDTO struct {
	Path       string `json:"path" bson:"path"`
	Key        string `json:"key" bson:"key"`
	Filename   string `json:"filename" bson:"filename"`
	User       string `json:"user" bson:"user"`
	FolderUUID string `json:"folder_uuid" bson:"folder_uuid"`
}

type FileUseCaseOutputDTO struct {
	Name       string    `json:"name" bson:"name"`
	User       string    `json:"user" bson:"user"`
	Key        string    `json:"key" bson:"key"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	Bucket     string    `json:"bucket" bson:"bucket"`
	SizeInMB   int64     `json:"size_mb" bson:"size_mb"`
	Size       int64     `json:"size" bson:"size"`
	FolderUUID string    `json:"folder_uuid" bson:"folder_uuid"`
}

type FileUseCaseGetInputDTO struct {
	Key string `json:"key" bson:"key"`
}

type FileUseCaseGetOutputDTO struct {
	FilePath  string    `json:"file_path" bson:"file_path"`
	File      *os.File  `json:"file" bson:"file"`
	Name      string    `json:"name" bson:"name"`
	User      string    `json:"user" bson:"user"`
	Key       string    `json:"key" bson:"key"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Bucket    string    `json:"bucket" bson:"bucket"`
	SizeInMB  int64     `json:"size_mb" bson:"size_mb"`
	Size      int64     `json:"size" bson:"size"`
}

type FileRepositoryInputDTO struct {
	Name       string    `json:"name" bson:"name"`
	User       string    `json:"user" bson:"user"`
	Key        string    `json:"key" bson:"key"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	Bucket     string    `json:"bucket" bson:"bucket"`
	SizeInMB   int64     `json:"size_mb" bson:"size_mb"`
	Size       int64     `json:"size" bson:"size"`
	FolderUUID string    `json:"folder_uuid" bson:"folder_uuid"`
}

type FileRepositoryOutputDTO struct {
	Name       string    `json:"name" bson:"name"`
	User       string    `json:"user" bson:"user"`
	Key        string    `json:"key" bson:"key"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	Bucket     string    `json:"bucket" bson:"bucket"`
	SizeInMB   int64     `json:"size_mb" bson:"size_mb"`
	Size       int64     `json:"size" bson:"size"`
	FolderUUID string    `json:"folder_uuid" bson:"folder_uuid"`
}

type FileRepositoryKeyInputDTO struct {
	Key string `json:"key" bson:"key"`
}

type FileRepositoryListAllInputDTO struct {
	FolderUUID string `json:"folder_uuid" bson:"folder_uuid"`
}

type FileUseCaseListAllInputDTO struct {
	FolderUUID string `json:"folder_uuid" bson:"folder_uuid"`
}

type FileStoreBucketInputDTO struct {
	Path     string `json:"path" bson:"path"`
	Key      string `json:"key" bson:"key"`
	Filename string `json:"filename" bson:"filename"`
	Bucket   string `json:"bucket_name" bson:"bucket_name"`
}

type FileStoreBucketOutputDTO struct {
	File     *os.File `json:"file" bson:"file"`
	Filename string   `json:"filename" bson:"filename"`
	FilePath string   `json:"file_path" bson:"file_path"`
	Key      string   `json:"key" bson:"key"`
}

type FileStoreBucketDeleteInputDTO struct {
	Key    string `json:"key" bson:"key"`
	Bucket string `json:"bucket_name" bson:"bucket_name"`
}
