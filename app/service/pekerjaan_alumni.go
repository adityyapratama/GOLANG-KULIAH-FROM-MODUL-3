package service

import (
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/app/repository"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

// GET /api/pekerjaan
func GetAllPekerjaan(c *fiber.Ctx) error {
    ctx := c.Context()
    list, err := repository.GetAllPekerjaan(ctx)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pekerjaan"})
    }
    return c.JSON(fiber.Map{"success": true, "data": list})
}

// GET /api/pekerjaan/:id
func GetPekerjaanByID(c *fiber.Ctx) error {
    ctx := c.Context()
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    }

    p, err := repository.GetPekerjaanByID(ctx, id)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
    }

    return c.JSON(fiber.Map{"success": true, "data": p})
}

// GET /api/pekerjaan/alumni/:alumni_id
func GetPekerjaanByAlumniID(c *fiber.Ctx) error {
    ctx := c.Context()
    alumniID, err := strconv.Atoi(c.Params("alumni_id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Alumni ID tidak valid"})
    }

    list, err := repository.GetPekerjaanByID(ctx, alumniID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pekerjaan"})
    }

    return c.JSON(fiber.Map{"success": true, "data": list})
}

// POST /api/pekerjaan
func CreatePekerjaan(c *fiber.Ctx) error {
    ctx := c.Context()
    var req model.CreatePekerjaanRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
    }

    // Convert ke model.Pekerjaan
    tanggalMulai, err := time.Parse("2006-01-02", req.TanggalMulaiKerja)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_mulai_kerja harus YYYY-MM-DD"})
    }

    var tanggalSelesai *time.Time
    if req.TanggalSelesaiKerja != nil {
        t, err := time.Parse("2006-01-02", *req.TanggalSelesaiKerja)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_selesai_kerja harus YYYY-MM-DD"})
        }
        tanggalSelesai = &t
    }

    pekerjaan := model.Pekerjaan{
        AlumniID:            req.AlumniID,
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

    if err := repository.CreatePekerjaan(ctx, &pekerjaan); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal menambah pekerjaan"})
    }

    pekerjaan.CreatedAt = time.Now()
    pekerjaan.UpdatedAt = time.Now()

    return c.Status(201).JSON(fiber.Map{"success": true, "data": pekerjaan})
}

// PUT /api/pekerjaan/:id
func UpdatePekerjaan(c *fiber.Ctx) error {
    ctx := c.Context()
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    }

    var req model.UpdatePekerjaanRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
    }

    tanggalMulai, err := time.Parse("2006-01-02", req.TanggalMulaiKerja)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_mulai_kerja harus YYYY-MM-DD"})
    }

    var tanggalSelesai *time.Time
    if req.TanggalSelesaiKerja != nil {
        t, err := time.Parse("2006-01-02", *req.TanggalSelesaiKerja)
        if err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Format tanggal_selesai_kerja harus YYYY-MM-DD"})
        }
        tanggalSelesai = &t
    }

    pekerjaan := model.Pekerjaan{
        ID:                  id,
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

    rows, err := repository.UpdatePekerjaan(ctx, &pekerjaan)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal update pekerjaan"})
    }
    if rows == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
    }

    pekerjaan.UpdatedAt = time.Now()

    return c.JSON(fiber.Map{"success": true, "data": pekerjaan})
}

// DELETE /api/pekerjaan/:id
func DeletePekerjaan(c *fiber.Ctx) error {
    ctx := c.Context()
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    }

    rows, err := repository.DeletePekerjaan(ctx, id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal hapus pekerjaan"})
    }
    if rows == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
    }

    return c.JSON(fiber.Map{"success": true, "message": "Pekerjaan berhasil dihapus"})
}
