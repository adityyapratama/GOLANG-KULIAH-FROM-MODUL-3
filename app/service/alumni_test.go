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

func TestCreateAlumni(t *testing.T) {
	app := fiber.New()
	mockRepo := new(MockAlumniRepo) 
	
	service := NewAlumniService(mockRepo)
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user_id", primitive.NewObjectID().Hex()) 
		return c.Next()
	})

	app.Post("/alumni", service.CreateAlumni)

	t.Run("Sukses - Data Valid & User Belum Ada", func(t *testing.T) {
		// Reset mock calls agar bersih
		mockRepo.ExpectedCalls = nil

		// Scenario:
		// 1. GetAlumniByUserID dipanggil -> Return Error NoDocuments (Artinya belum ada, boleh lanjut)
		// 2. CreateAlumni dipanggil -> Return nil (Sukses)

		mockRepo.On("GetAlumniByUserID", mock.Anything, mock.Anything).
			Return(nil, mongo.ErrNoDocuments).Once()

		mockRepo.On("CreateAlumni", mock.Anything, mock.Anything).
			Return(nil).Once()

		payload := model.CreateAlumniRequest{
			NIM: "12345", Nama: "Budi", Email: "budi@mail.com",
			Jurusan: "IT", Angkatan: 2020, TahunLulus: 2024,
		}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/alumni", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, 201, resp.StatusCode)
	})

	t.Run("Gagal - User Sudah Terdaftar (Conflict)", func(t *testing.T) {
		// Reset mock calls
		mockRepo.ExpectedCalls = nil

		// Scenario:
		// 1. GetAlumniByUserID dipanggil -> Return Data Alumni (Artinya sudah ada)
		// 2. CreateAlumni TIDAK BOLEH dipanggil

		existingAlumni := &model.Alumni{Nama: "Sudah Ada"}
		
		mockRepo.On("GetAlumniByUserID", mock.Anything, mock.Anything).
			Return(existingAlumni, nil).Once() // Nil error means found

		payload := model.CreateAlumniRequest{NIM: "123", Nama: "Baru", Email: "baru@mail.com"}
		body, _ := json.Marshal(payload)
		req := httptest.NewRequest("POST", "/alumni", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")

		resp, _ := app.Test(req)
		assert.Equal(t, 409, resp.StatusCode) // Conflict
	})
}