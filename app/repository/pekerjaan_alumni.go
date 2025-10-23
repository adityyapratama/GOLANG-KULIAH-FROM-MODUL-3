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



type IPekerjaanRepository interface {
	CreatePekerjaan(ctx context.Context, pekerjaan *model.Pekerjaan) error
	GetPekerjaanByID(ctx context.Context, id string) (*model.Pekerjaan, error)
	GetPekerjaanByAlumniID(ctx context.Context, alumniID primitive.ObjectID) ([]model.Pekerjaan, error)
	GetAllPekerjaan(ctx context.Context) ([]model.Pekerjaan, error)
	UpdatePekerjaan(ctx context.Context, id string, pekerjaan *model.Pekerjaan) (int64, error)
	DeletePekerjaan(ctx context.Context, id string) (int64, error)
}


type pekerjaanRepository struct {
	collection *mongo.Collection
}



func NewPekerjaanRepository(db *mongo.Database) IPekerjaanRepository {
	return &pekerjaanRepository{
		collection: db.Collection("pekerjaan_alumni"),
	}
}



func (r *pekerjaanRepository) CreatePekerjaan(ctx context.Context, pekerjaan *model.Pekerjaan) error {
	pekerjaan.CreatedAt = time.Now()
	pekerjaan.UpdatedAt = time.Now()

	result, err := r.collection.InsertOne(ctx, pekerjaan)
	if err != nil {
		return err
	}
	pekerjaan.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *pekerjaanRepository) GetPekerjaanByID(ctx context.Context, id string) (*model.Pekerjaan, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var pekerjaan model.Pekerjaan
	filter := bson.M{"_id": objID, "deleted_at": nil}

	if err := r.collection.FindOne(ctx, filter).Decode(&pekerjaan); err != nil {
		return nil, err
	}
	return &pekerjaan, nil
}

func (r *pekerjaanRepository) GetPekerjaanByAlumniID(ctx context.Context, alumniID primitive.ObjectID) ([]model.Pekerjaan, error) {
	var pekerjaanList []model.Pekerjaan
	filter := bson.M{"alumni_id": alumniID, "deleted_at": nil}
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &pekerjaanList); err != nil {
		return nil, err
	}
	return pekerjaanList, nil
}

func (r *pekerjaanRepository) GetAllPekerjaan(ctx context.Context) ([]model.Pekerjaan, error) {
	var pekerjaanList []model.Pekerjaan
	filter := bson.M{"deleted_at": nil}
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err = cursor.All(ctx, &pekerjaanList); err != nil {
		return nil, err
	}
	return pekerjaanList, nil
}

func (r *pekerjaanRepository) UpdatePekerjaan(ctx context.Context, id string, pekerjaan *model.Pekerjaan) (int64, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": objID, "deleted_at": nil}
	
	update := bson.M{
		"$set": bson.M{
			"nama_perusahaan":       pekerjaan.NamaPerusahaan,
			"posisi_jabatan":        pekerjaan.PosisiJabatan,
			"bidang_industri":       pekerjaan.BidangIndustri,
			"lokasi_kerja":          pekerjaan.LokasiKerja,
			"gaji_range":            pekerjaan.GajiRange,
			"tanggal_mulai_kerja":   pekerjaan.TanggalMulaiKerja,
			"tanggal_selesai_kerja": pekerjaan.TanggalSelesaiKerja,
			"status_pekerjaan":      pekerjaan.StatusPekerjaan,
			"deskripsi_pekerjaan":   pekerjaan.DeskripsiPekerjaan,
			"updated_at":            time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (r *pekerjaanRepository) DeletePekerjaan(ctx context.Context, id string) (int64, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return 0, err
	}

	filter := bson.M{"_id": objID, "deleted_at": nil}
	
	// Soft Delete
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}