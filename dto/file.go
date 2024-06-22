package dto

import (
	"os"
	"time"
)

type FileUseCaseInputDTO struct {
	Path     string `json:"path" bson:"path"`
	Filename string `json:"filename" bson:"filename"`
	User     string `json:"user" bson:"user"`
}

type FileUseCaseOutputDTO struct {
	Name      string    `json:"name" bson:"name"`
	User      string    `json:"user" bson:"user"`
	Key       string    `json:"key" bson:"key"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Bucket    string    `json:"bucket" bson:"bucket"`
	SizeInMB  int64     `json:"size_mb" bson:"size_mb"`
	Size      int64     `json:"size" bson:"size"`
}

type FileUseCaseGetInputDTO struct {
	Key string `json:"key" bson:"key"`
}

type FileUseCaseGetOutputDTO struct {
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
	Name      string    `json:"name" bson:"name"`
	User      string    `json:"user" bson:"user"`
	Key       string    `json:"key" bson:"key"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Bucket    string    `json:"bucket" bson:"bucket"`
	SizeInMB  int64     `json:"size_mb" bson:"size_mb"`
	Size      int64     `json:"size" bson:"size"`
}

type FileRepositoryOutputDTO struct {
	Name      string    `json:"name" bson:"name"`
	User      string    `json:"user" bson:"user"`
	Key       string    `json:"key" bson:"key"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	Bucket    string    `json:"bucket" bson:"bucket"`
	SizeInMB  int64     `json:"size_mb" bson:"size_mb"`
	Size      int64     `json:"size" bson:"size"`
}

type FileRepositoryKeyInputDTO struct {
	Key string `json:"key" bson:"key"`
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
	Key      string   `json:"key" bson:"key"`
}
