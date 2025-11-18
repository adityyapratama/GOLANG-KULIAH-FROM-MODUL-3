package service

import (
	"context"
	"golang-kuliah-from-modul-3/app/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type MockAlumniRepo struct {
	mock.Mock
}

func (m *MockAlumniRepo) GetAlumniByUserID(ctx context.Context, userID string) (*model.Alumni, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Alumni), args.Error(1)
}

func (m *MockAlumniRepo) CreateAlumni(ctx context.Context, alumni *model.Alumni) error {
	args := m.Called(ctx, alumni)
	return args.Error(0)
}

func (m *MockAlumniRepo) GetAlumniByID(ctx context.Context, id string) (*model.Alumni, error) {
	args := m.Called(ctx, id)
	// PENTING: Pastikan return nil jika args(0) nil agar error terbaca
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Alumni), args.Error(1)
}


func (m *MockAlumniRepo) GetAllAlumni(ctx context.Context) ([]model.Alumni, error) { return nil, nil }
func (m *MockAlumniRepo) UpdateAlumni(ctx context.Context, id string, alumni *model.Alumni) (int64, error) { return 0, nil }
func (m *MockAlumniRepo) DeleteAlumni(ctx context.Context, id string) (int64, error) { return 0, nil }






type MockPekerjaanRepo struct {
	mock.Mock
}

func (m *MockPekerjaanRepo) CreatePekerjaan(ctx context.Context, p *model.Pekerjaan) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

// Dummy Methods
func (m *MockPekerjaanRepo) GetPekerjaanByID(ctx context.Context, id string) (*model.Pekerjaan, error) { return nil, nil }
func (m *MockPekerjaanRepo) GetPekerjaanByAlumniID(ctx context.Context, id primitive.ObjectID) ([]model.Pekerjaan, error) { return nil, nil }
func (m *MockPekerjaanRepo) GetAllPekerjaan(ctx context.Context) ([]model.Pekerjaan, error) { return nil, nil }
func (m *MockPekerjaanRepo) UpdatePekerjaan(ctx context.Context, id string, p *model.Pekerjaan) (int64, error) { return 0, nil }
func (m *MockPekerjaanRepo) DeletePekerjaan(ctx context.Context, id string) (int64, error) { return 0, nil }






type MockFileRepo struct {
	mock.Mock
}

func (m *MockFileRepo) Create(file *model.File) error {
	args := m.Called(file)
	return args.Error(0)
}

// Dummy Methods
func (m *MockFileRepo) FindAll() ([]model.File, error) { return nil, nil }
func (m *MockFileRepo) FindByID(id string) (*model.File, error) { return nil, nil }
func (m *MockFileRepo) Delete(id string) error { return nil }
func (m *MockFileRepo) DeleteByUser(fileID, userID string) (int64, error) { return 0, nil }