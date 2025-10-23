package repository

import (
	"context"
	"golang-kuliah-from-modul-3/app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type IAlumniRepository interface {
    CreateAlumni(ctx context.Context, alumni *model.Alumni) error
    GetAlumniByID(ctx context.Context,id string) (*model.Alumni, error) 
    GetAlumniByUserID(ctx context.Context, userID string) (*model.Alumni, error)
	GetAllAlumni(ctx context.Context) ([]model.Alumni, error)
	UpdateAlumni(ctx context.Context, id string, alumni *model.Alumni) (int64, error)
	DeleteAlumni(ctx context.Context, id string) (int64, error)
}

type alumniRepository struct{
    collection  *mongo.Collection
}

func NewAlumniRepository(db *mongo.Database) IAlumniRepository {
    return &alumniRepository{
        collection: db.Collection("alumni"),
    }
} 

func (r *alumniRepository)CreateAlumni(ctx context.Context, alumni *model.Alumni) error{
    alumni.CreatedAt = time.Now()
    result, err := r.collection.InsertOne(ctx, alumni)
    if err != nil{
        return err
    }
    alumni.ID = result.InsertedID.(primitive.ObjectID)

    return  nil
    
}


func (r *alumniRepository)GetAlumniByID(ctx context.Context, id string) (*model.Alumni, error){
    objID, err := primitive.ObjectIDFromHex(id)

    if err != nil{
        return  nil , err
    }
    var alumni model.Alumni
    filter := bson.M{"_id":objID, "deleted_at": nil}
    if err := r.collection.FindOne(ctx, filter).Decode(&alumni); err != nil{
        return nil, err
    }

    return  &alumni, nil

}


func (r *alumniRepository) GetAlumniByUserID(ctx context.Context, userID string) (*model.Alumni, error) {
	userObjID, err := primitive.ObjectIDFromHex(userID)

	if err != nil {
		return nil, err
	}

	var alumni model.Alumni
	filter := bson.M{"user_id": userObjID, "deleted_at": nil}

	if err := r.collection.FindOne(ctx, filter).Decode(&alumni); err != nil {
		return nil, err
	}
	return &alumni, nil
}


func (r *alumniRepository)GetAllAlumni(ctx context.Context) ([]model.Alumni, error){
    var lisAlumni []model.Alumni
    filter := bson.M{"deleted_at": nil}

    opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

    cursor , err :=r.collection.Find(ctx, filter, opts)
    if err != nil{
        return nil, err
    }

    defer cursor.Close(ctx)

    if err =cursor.All(ctx, &lisAlumni); err != nil{
        return nil, err
    }
    
    return lisAlumni, nil


}

func (r *alumniRepository)UpdateAlumni(ctx context.Context, id string, alumni *model.Alumni) (int64, error){
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil{
        return  0, err
    }

    filter  :=bson.M{"_id":objID, "deleted_at": nil}
    update:=bson.M{
        "$set":bson.M{
            "nim":         alumni.NIM,
			"nama":        alumni.Nama,
			"jurusan":     alumni.Jurusan,
			"angkatan":    alumni.Angkatan,
			"tahun_lulus": alumni.TahunLulus,
			"email":       alumni.Email,
			"no_telepon":  alumni.NoTelepon,
			"alamat":      alumni.Alamat,
			"updated_at":  time.Now(),
            
        },
    }
        result, err := r.collection.UpdateOne(ctx, filter, update)
        if err != nil{
            return 0, err
        }
        
        return result.ModifiedCount, nil
    }


    func (r *alumniRepository)DeleteAlumni(ctx context.Context, id string) (int64, error){
        objID, err:= primitive.ObjectIDFromHex(id)
        if err !=nil{
            return 0, err
        }

        filter :=bson.M{"_id": objID, "deleted_at": nil}
        update :=bson.M{
            "$set":bson.M{
                "deleted_at": time.Now(),
            },
        }

        result, err := r.collection.UpdateOne(ctx, filter, update)
        if err != nil{
            return 0, err
        }

        return result.ModifiedCount, nil
    }



