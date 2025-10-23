package service

import (
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/app/repository"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPekerjaanService interface {
	CreatePekerjaan(c *fiber.Ctx) error
	GetPekerjaanByID(c *fiber.Ctx) error
	GetPekerjaanByAlumniID(c *fiber.Ctx) error
	GetAllPekerjaan(c *fiber.Ctx) error
	UpdatePekerjaan(c *fiber.Ctx) error
	DeletePekerjaan(c *fiber.Ctx) error
}

type pekerjaanService struct {
	pekerjaanRepo repository.IPekerjaanRepository
	alumniRepo    repository.IAlumniRepository // Dibutuhkan untuk memvalidasi AlumniID
}


func NewPekerjaanService(
	pekerjaanRepo repository.IPekerjaanRepository,
	alumniRepo repository.IAlumniRepository,
) IPekerjaanService {
	return &pekerjaanService{
		pekerjaanRepo: pekerjaanRepo,
		alumniRepo:    alumniRepo,
	}
}



func parseTanggal(tgl string) (time.Time, error) {
	// Format date
	return time.Parse("2006-01-02", tgl)
}




func (s *pekerjaanService) CreatePekerjaan(c *fiber.Ctx) error {
	var req model.CreatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
	}

	// Validasi data
	if req.AlumniID == "" || req.NamaPerusahaan == "" || req.PosisiJabatan == "" || req.TanggalMulaiKerja == "" {
		return c.Status(400).JSON(fiber.Map{"error": "AlumniID, NamaPerusahaan, PosisiJabatan, dan TanggalMulaiKerja wajib diisi"})
	}

	
	alumniObjID, err := primitive.ObjectIDFromHex(req.AlumniID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format AlumniID tidak valid"})
	}

	
	_, err = s.alumniRepo.GetAlumniByID(c.Context(), req.AlumniID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"error": "Alumni dengan ID tersebut tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengecek data alumni"})
	}

	
	tanggalMulai, err := parseTanggal(req.TanggalMulaiKerja)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_mulai_kerja harus YYYY-MM-DD"})
	}

	var tanggalSelesai *time.Time
	if req.TanggalSelesaiKerja != nil {
		t, err := parseTanggal(*req.TanggalSelesaiKerja)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_selesai_kerja harus YYYY-MM-DD"})
		}
		tanggalSelesai = &t
	}

	pekerjaan := model.Pekerjaan{
		AlumniID:            alumniObjID,
		NamaPerusahaan:      req.NamaPerusahaan,
		PosisiJabatan:       req.PosisiJabatan,
		BidangIndustri:      req.BidangIndustri,
		LokasiKerja:         req.LokasiKerja,
		GajiRange:           req.GajiRange,
		TanggalMulaiKerja:   tanggalMulai,
		TanggalSelesaiKerja: tanggalSelesai,
		StatusPekerjaan:     req.StatusPekerjaan,
		DeskripsiPekerjaan:  req.DeskripsiPekerjaan,
	}

	if err := s.pekerjaanRepo.CreatePekerjaan(c.Context(), &pekerjaan); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal menyimpan data pekerjaan"})
	}

	return c.Status(201).JSON(fiber.Map{"success": true, "data": pekerjaan})
}


func (s *pekerjaanService) GetPekerjaanByID(c *fiber.Ctx) error {
	id := c.Params("id")
	pekerjaan, err := s.pekerjaanRepo.GetPekerjaanByID(c.Context(), id)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data"})
	}
	return c.JSON(fiber.Map{"success": true, "data": pekerjaan})
}


func (s *pekerjaanService) GetPekerjaanByAlumniID(c *fiber.Ctx) error {
	alumniID := c.Params("alumni_id")
	alumniObjID, err := primitive.ObjectIDFromHex(alumniID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format AlumniID tidak valid"})
	}

	list, err := s.pekerjaanRepo.GetPekerjaanByAlumniID(c.Context(), alumniObjID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pekerjaan"})
	}
	return c.JSON(fiber.Map{"success": true, "data": list})
}

func (s *pekerjaanService) GetAllPekerjaan(c *fiber.Ctx) error {
	list, err := s.pekerjaanRepo.GetAllPekerjaan(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pekerjaan"})
	}
	return c.JSON(fiber.Map{"success": true, "data": list})
}


func (s *pekerjaanService) UpdatePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	var req model.UpdatePekerjaanRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
	}

	// Parse Tanggal
	tanggalMulai, err := parseTanggal(req.TanggalMulaiKerja)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_mulai_kerja harus YYYY-MM-DD"})
	}

	var tanggalSelesai *time.Time
	if req.TanggalSelesaiKerja != nil {
		t, err := parseTanggal(*req.TanggalSelesaiKerja)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_selesai_kerja harus YYYY-MM-DD"})
		}
		tanggalSelesai = &t
	}

	pekerjaan := model.Pekerjaan{
		NamaPerusahaan:      req.NamaPerusahaan,
		PosisiJabatan:       req.PosisiJabatan,
		BidangIndustri:      req.BidangIndustri,
		LokasiKerja:         req.LokasiKerja,
		GajiRange:           req.GajiRange,
		TanggalMulaiKerja:   tanggalMulai,
		TanggalSelesaiKerja: tanggalSelesai,
		StatusPekerjaan:     req.StatusPekerjaan,
		DeskripsiPekerjaan:  req.DeskripsiPekerjaan,
	}

	rows, err := s.pekerjaanRepo.UpdatePekerjaan(c.Context(), id, &pekerjaan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal update pekerjaan"})
	}
	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
	}

	return c.JSON(fiber.Map{"success": true, "message": "Pekerjaan berhasil diupdate"})
}

func (s *pekerjaanService) DeletePekerjaan(c *fiber.Ctx) error {
	id := c.Params("id")
	rows, err := s.pekerjaanRepo.DeletePekerjaan(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal hapus pekerjaan"})
	}
	if rows == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
	}
	return c.JSON(fiber.Map{"success": true, "message": "Pekerjaan berhasil dihapus"})
}