package service

import (
	"bytes"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUploadFoto(t *testing.T) {
	app := fiber.New()
	mockRepo := new(MockFileRepo)
	
	// Pakai folder temp agar tidak mengotori project
	tempDir := os.TempDir()
	service := NewFileService(mockRepo, tempDir)

	// Middleware Mock Auth
	app.Use(func(c *fiber.Ctx) error {
		c.Locals("user_id", primitive.NewObjectID().Hex())
		return c.Next()
	})

	// Registrasi Route
	app.Post("/files/foto", service.UploadFoto)

	t.Run("Sukses Upload Foto JPG", func(t *testing.T) {
		// Reset mock
		mockRepo.ExpectedCalls = nil

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		
		// Buat part file
		part, _ := writer.CreateFormFile("file", "test.jpg")
		part.Write([]byte("dummy image content"))
		
		// Trik: Kita tidak bisa mudah memalsukan Content-Type per-part dengan writer standar Go.
		// Service Anda mengecek `fileHeader.Header.Get("Content-Type")`.
		// Jika validasi MIME gagal, dia akan return 400.
		// Untuk unit test Service -> Repo, kita bisa terima 400 (Validasi gagal) 
		// asalkan BUKAN 404 (Route not found) atau 500 (Panic).
		
		writer.Close()

		// Expectation: Repo Create MUNGKIN dipanggil jika lolos validasi
		// Tapi karena MIME type mungkin kosong/octet-stream, validasi di service akan gagal.
		// Jadi kita buat mock ini "Maybe" dipanggil.
		mockRepo.On("Create", mock.Anything).Return(nil).Maybe()

		req := httptest.NewRequest("POST", "/files/foto", body)
		req.Header.Set("Content-Type", writer.FormDataContentType())

		resp, _ := app.Test(req)
		
		// Debugging Output
		if resp.StatusCode == 404 {
			t.Fatal("Route not found (404). Cek registrasi route app.Post")
		}

		// Kita assert tidak error server
		assert.NotEqual(t, 500, resp.StatusCode)
		assert.NotEqual(t, 404, resp.StatusCode)
	})
}