package dto

import "time"

type FolderUseCaseCreateInputDTO struct {
	Name string `json:"name" bson:"name"`
	User string `json:"user" bson:"user"`
}

type FolderUseCaseCreateOutputDTO struct {
	UUID string `json:"uuid" bson:"uuid"`
	Name string `json:"name" bson:"name"`
}

type FolderRepositoryInputDTO struct {
	UUID      string    `json:"uuid" bson:"uuid"`
	Name      string    `json:"name" bson:"name"`
	User      string    `json:"user" bson:"user"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

type FolderRepositoryOutputDTO struct {
	UUID string `json:"uuid" bson:"uuid"`
	Name string `json:"name" bson:"name"`
}
