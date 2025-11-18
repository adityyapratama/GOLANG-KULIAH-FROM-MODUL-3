package repository

import (
	"context"
	"golang-kuliah-from-modul-3/app/model"
)

type AuthRepository interface {
	UserLogin(ctx context.Context, userLogin string) (*model.User, string, error)
	CreateUser(ctx context.Context, user *model.User, passwordHash string) error
}


type AlumniRepository interface {
	GetAllAlumni (ctx context.Context) ([]model.Alumni, error)
	GetAlumniByID(ctx context.Context, id int) (*model.Alumni, error)
	CreateAlumni(ctx context.Context ,a *model.Alumni) error
	UpdateAlumni(ctx context.Context, a *model.Alumni) (int64 , error)
	DeleteAlumni(ctx context.Context, id int) (int64, error)
}

type PekerjaanRepository interface{
	GetAllPekerjaan (ctx context.Context) ([]model.Pekerjaan, error)
	GetPekerjaanByID(ctx context.Context, id int) (*model.Pekerjaan, error)
	CreatePekerjaani(ctx context.Context ,a *model.Pekerjaan) error
	UpdatePekerjaan(ctx context.Context, a *model.Pekerjaan) (int64 , error)
	DeletePekerjaan(ctx context.Context, id int) (int64, error)

}