package repository

import (
	"context"
	"dropstore/dto"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FolderRepository struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewFolderRepository(ctx context.Context, collection *mongo.Collection) *FolderRepository {
	return &FolderRepository{ctx, collection}
}

func (r *FolderRepository) Save(input dto.FolderRepositoryInputDTO) (dto.FolderRepositoryOutputDTO, error) {
	var result dto.FolderRepositoryOutputDTO
	input.CreatedAt = time.Now()
	res, err := r.collection.InsertOne(r.ctx, &input)

	if err != nil {
		fmt.Printf("Error inserting folder in DB, error: %v", err)
		return result, err
	}

	query := bson.M{"_id": res.InsertedID}
	err = r.collection.FindOne(r.ctx, query).Decode(&result)
	if err != nil {
		fmt.Printf("Error getting folder in DB, error: %v", err)
		return result, err
	}
	return result, nil
}

func (r *FolderRepository) ListAll() ([]dto.FolderRepositoryOutputDTO, error) {
	var folders []dto.FolderRepositoryOutputDTO
	query := bson.M{}
	cur, err := r.collection.Find(r.ctx, query)
	if err != nil {
		fmt.Printf("Error listing folder in DB, error: %v", err)
		return nil, err
	}
	for cur.Next(r.ctx) {
		var folder dto.FolderRepositoryOutputDTO
		err = cur.Decode(&folder)
		if err != nil {
			fmt.Printf("Error decoding file from db, err %v", err)
			continue
		}
		folders = append(folders, folder)
	}
	return folders, nil
}
