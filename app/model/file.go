package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type File struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID       primitive.ObjectID `json:"user_id" bson:"user_id"`   
	Category     string             `json:"category" bson:"category"` 
	FileName     string             `json:"file_name" bson:"file_name"`
	OriginalName string             `json:"original_name" bson:"original_name"`
	FilePath     string             `json:"file_path" bson:"file_path"`
	FileSize     int64              `json:"file_size" bson:"file_size"`
	FileType     string             `json:"file_type" bson:"file_type"`
	UploadedAt   time.Time          `json:"uploaded_at" bson:"uploaded_at"`
	UploadedBy   primitive.ObjectID `json:"uploaded_by" bson:"uploaded_by"` 
}


type FileResponse struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	Category     string    `json:"category"`
	FileName     string    `json:"file_name"`
	OriginalName string    `json:"original_name"`
	FilePath     string    `json:"file_path"`
	FileSize     int64     `json:"file_size"`
	FileType     string    `json:"file_type"`
	UploadedAt   time.Time `json:"uploaded_at"`
	UploadedBy   string    `json:"uploaded_by"`
}
