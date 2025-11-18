package service

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"golang-kuliah-from-modul-3/app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestCreatePekerjaan(t *testing.T) {
	// Setup Fiber
	app := fiber.New()
	mockPekerjaanRepo := new(MockPekerjaanRepo)
	mockAlumniRepo := new(MockAlumniRepo)
	service := NewPekerjaanService(mockPekerjaanRepo, mockAlumniRepo)
	
	app.Post("/pekerjaan", service.CreatePekerjaan)

	t.Run("Sukses Tambah Pekerjaan", func(t *testing.T) {
		validAlumniID := primitive.NewObjectID().Hex()

		// Reset expects agar bersih
		mockAlumniRepo.ExpectedCalls = nil
		mockPekerjaanRepo.ExpectedCalls = nil

		// 1. Expect Alumni Ditemukan
		mockAlumniRepo.On("GetAlumniByID", mock.Anything, validAlumniID).
			Return(&model.Alumni{}, nil).Once()

		// 2. Expect Simpan Pekerjaan
		mockPekerjaanRepo.On("CreatePekerjaan", mock.Anything, mock.Anything).
			Return(nil).Once()

		payload := model.CreatePekerjaanRequest{
			AlumniID: validAlumniID,
			NamaPerusahaan: "Google",
			PosisiJabatan: "Engineer",
			TanggalMulaiKerja: "2023-01-01",
		}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/pekerjaan", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, 201, resp.StatusCode)
	})

	t.Run("Gagal - Alumni Tidak Ditemukan", func(t *testing.T) {
		missingID := primitive.NewObjectID().Hex()

		// Reset expects lagi
		mockAlumniRepo.ExpectedCalls = nil
		mockPekerjaanRepo.ExpectedCalls = nil

		// 1. Expect Alumni Gagal Ditemukan (Return Error Mongo)
		mockAlumniRepo.On("GetAlumniByID", mock.Anything, missingID).
			Return(nil, mongo.ErrNoDocuments).Once()

		// Kita TIDAK set expect CreatePekerjaan, karena kode harus berhenti di atas.

		payload := model.CreatePekerjaanRequest{
			AlumniID: missingID,
			NamaPerusahaan: "Google",
			PosisiJabatan: "Engineer",
			TanggalMulaiKerja: "2023-01-01",
		}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/pekerjaan", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		
		// Debugging jika status bukan 404
		if resp.StatusCode != 404 {
			t.Logf("Expected 404, got %d", resp.StatusCode)
		}
		assert.Equal(t, 404, resp.StatusCode)
	})
}