package repository

import (
	"context"
	"golang-kuliah-from-modul-3/app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FileRepository interface {
	Create(file *model.File) error
	FindAll() ([]model.File, error)
	FindByID(id string) (*model.File, error)
	Delete(id string) error
	DeleteByUser(fileID string, userID string) (int64, error) 
}

type NewFileRepository struct {
	collection *mongo.Collection
}

func NewRepositoryFile(db *mongo.Database) FileRepository {
	return &NewFileRepository{
		collection: db.Collection("files"),
	}
}

func (r *NewFileRepository) Create(file *model.File) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	file.UploadedAt = time.Now()
	result, err := r.collection.InsertOne(ctx, file)
	if err != nil {
		return err
	}

	file.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *NewFileRepository) FindAll() ([]model.File, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var files []model.File
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &files); err != nil {
		return nil, err
	}

	return files, err

}

func (r *NewFileRepository) FindByID(id string) (*model.File, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var file model.File
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&file)
	if err != nil {
		return nil, err
	}
	return &file, nil
}

func (r *NewFileRepository) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err

}


func (r *NewFileRepository) DeleteByUser(fileID string, userID string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert string ID ke ObjectID
	fileObjectID, err := primitive.ObjectIDFromHex(fileID)
	if err != nil {
		return 0, err
	}

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return 0, err
	}

	
	filter := bson.M{
		"_id":        fileObjectID,
		"created_by": userObjectID, 
	}

	result, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}

	
	return result.DeletedCount, nil
}
