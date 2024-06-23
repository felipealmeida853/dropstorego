package repository

import (
	"context"
	"dropstore/dto"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileRepository struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewFileRepository(ctx context.Context, collection *mongo.Collection) *FileRepository {
	return &FileRepository{ctx, collection}
}

func (r *FileRepository) Save(input dto.FileRepositoryInputDTO) (dto.FileRepositoryOutputDTO, error) {
	var result dto.FileRepositoryOutputDTO
	res, err := r.collection.InsertOne(r.ctx, &input)

	if err != nil {
		fmt.Printf("Error inserting file in DB, error: %v", err)
		return result, err
	}

	query := bson.M{"_id": res.InsertedID}
	err = r.collection.FindOne(r.ctx, query).Decode(&result)
	if err != nil {
		fmt.Printf("Error getting file in DB, error: %v", err)
		return result, err
	}
	return result, nil
}

func (r *FileRepository) GetByKey(input dto.FileRepositoryKeyInputDTO) (dto.FileRepositoryOutputDTO, error) {
	var result dto.FileRepositoryOutputDTO

	query := bson.M{"key": input.Key}
	err := r.collection.FindOne(r.ctx, query).Decode(&result)
	if err != nil {
		fmt.Printf("Error getting file in DB, error: %v", err)
		return result, err
	}

	return result, nil
}

func (r *FileRepository) DeleteByKey(input dto.FileRepositoryKeyInputDTO) error {
	query := bson.M{"key": input.Key}
	_, err := r.collection.DeleteOne(r.ctx, query)
	if err != nil {
		fmt.Printf("Error deleting file in DB, error: %v", err)
	}
	return err
}

func (r *FileRepository) ListAll() ([]dto.FileRepositoryOutputDTO, error) {
	var files []dto.FileRepositoryOutputDTO
	query := bson.M{}
	cur, err := r.collection.Find(r.ctx, query)
	if err != nil {
		fmt.Printf("Error listing file in DB, error: %v", err)
		return nil, err
	}
	for cur.Next(r.ctx) {
		var file dto.FileRepositoryOutputDTO
		err = cur.Decode(&file)
		if err != nil {
			fmt.Printf("Error decoding file from db, err %v", err)
			continue
		}
		files = append(files, file)
	}
	return files, nil
}
