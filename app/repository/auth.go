package repository

import (
	"context"
	"golang-kuliah-from-modul-3/app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAuthRepository interface {
    GetUserByLogin(ctx context.Context, login string) (*model.User , error)
    CreateUser(ctx context.Context, User *model.User) error

}

type authRepository struct{
    userCollection  *mongo.Collection
}

func NewAuthRepository(db *mongo.Database) IAuthRepository {
	return &authRepository{
		userCollection:      db.Collection("users"),
	}
}


func (r *authRepository) GetUserByLogin(ctx context.Context, login string) (*model.User , error) {
    var user model.User

    filter := bson.M{
        "$or" :[]bson.M{
            {"username": login},
            {"email": login},
        },
    }
    
    err := r.userCollection.FindOne(ctx, filter).Decode(&user)
    if err != nil{
        if err == mongo.ErrNoDocuments{
            return nil, nil
        }

        return nil, err
    }

    return &user, nil

}

func (r *authRepository) CreateUser(ctx context.Context, user *model.User) error {
    user.CreatedAt = time.Now()

    result, err := r.userCollection.InsertOne(ctx,user)
    if err != nil{
        return  err
    }

    user.ID = result.InsertedID.(primitive.ObjectID)

    return  nil

}