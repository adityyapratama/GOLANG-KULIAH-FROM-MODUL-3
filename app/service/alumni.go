package service

import (
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/app/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IAlumniService interface {
	CreateAlumni(c *fiber.Ctx) error
	GetAlumniByID(c *fiber.Ctx) error
	GetAlumniByUserID(c *fiber.Ctx) error // <-- Metode ini harus diimplementasi
	GetAllAlumni(c *fiber.Ctx) error
	UpdateAlumni(c *fiber.Ctx) error
	DeleteAlumni(c *fiber.Ctx) error
}

type alumniService struct {
	alumniRepo repository.IAlumniRepository
	// pekerjaanRepo repository.IPekerjaanRepository
}

func NewAlumniService(
	alumniRepo repository.IAlumniRepository,
	// pekerjaanRepo repository.IPekerjaanRepository,
) IAlumniService {
	return &alumniService{
		alumniRepo: alumniRepo,
		// pekerjaanRepo: pekerjaanRepo,
	}
}

func (s *alumniService) CreateAlumni(c *fiber.Ctx) error {
	var req model.CreateAlumniRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	userIDHex := c.Locals("user_id").(string)
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse user ID",
		})
	}

	if req.Nama == "" || req.NIM == "" || req.Email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "NIM, Nama, dan Email wajib diisi"})
	}

	// Logika ini sudah BENAR
	_, err = s.alumniRepo.GetAlumniByUserID(c.Context(), userIDHex)
	if err == nil {
		// Alumni sudah ada, kirim 409 Conflict
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "Alumni dengan user ID tersebut sudah ada"})
	} else if err != mongo.ErrNoDocuments {
		// Terjadi error database LAINNYA, jangan lanjut
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Gagal mengecek data alumni"})
	}

	alumni := model.Alumni{
		UserID:     userID, // Set UserID dari token
		NIM:        req.NIM,
		Nama:       req.Nama,
		Jurusan:    req.Jurusan,
		Angkatan:   req.Angkatan,
		TahunLulus: req.TahunLulus,
		Email:      req.Email,
		NoTelepon:  req.NoTelepon,
		Alamat:     req.Alamat,
	}

	if err := s.alumniRepo.CreateAlumni(c.Context(), &alumni); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal membuat data alumni"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Alumni created successfully",
		"alumni":  alumni,
	})
}

func (s *alumniService) GetAlumniByID(c *fiber.Ctx) error {
	id := c.Params("id")

	alumni, err := s.alumniRepo.GetAlumniByID(c.Context(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}

	return c.JSON(fiber.Map{"success": true, "data": alumni})
}

func (s *alumniService) GetAllAlumni(c *fiber.Ctx) error {
	list, err := s.alumniRepo.GetAllAlumni(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data alumni"})
	}
	return c.JSON(fiber.Map{"success": true, "data": list})
}

func (s *alumniService) UpdateAlumni(c *fiber.Ctx) error {
	id := c.Params("id")
	var req model.UpdateAlumniRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
	}

	alumniToUpdate := &model.Alumni{
		NIM:        req.NIM,
		Nama:       req.Nama,
		Jurusan:    req.Jurusan,
		Angkatan:   req.Angkatan,
		TahunLulus: req.TahunLulus,
		Email:      req.Email,
		NoTelepon:  req.NoTelepon,
		Alamat:     req.Alamat,
	}

	rows, err := s.alumniRepo.UpdateAlumni(c.Context(), id, alumniToUpdate)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal memperbarui data alumni"})
	}

	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "alumni berhasil di update"})
}

func (s *alumniService) DeleteAlumni(c *fiber.Ctx) error {
	// PERBAIKAN 1: Tambahkan ':='
	id := c.Params("id")

	rows, err := s.alumniRepo.DeleteAlumni(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menghapus data alumni"})
	}

	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Alumni berhasil dihapus",
	})
}

// PERBAIKAN 2: Tambahkan implementasi GetAlumniByUserID yang hilang
func (s *alumniService) GetAlumniByUserID(c *fiber.Ctx) error {
	// Asumsi kita ambil 'userid' dari URL param (parameter rute)
	userID := c.Params("userid")
	if userID == "" {
		return c.Status(400).JSON(fiber.Map{"error": "UserID param tidak boleh kosong"})
	}

	alumni, err := s.alumniRepo.GetAlumniByUserID(c.Context(), userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"error": "Alumni dengan UserID tersebut tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}

	return c.JSON(fiber.Map{"success": true, "data": alumni})
}