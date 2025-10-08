package service

import (
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/app/repository"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)


func GetAllPekerjaan(c *fiber.Ctx) error {
    ctx := c.Context()
    list, err := repository.GetAllPekerjaan(ctx)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pekerjaan"})
    }
    return c.JSON(fiber.Map{"success": true, "data": list})
}


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

func GetTotalKerja(c *fiber.Ctx) error {
    ctx := c.Context()
    alumniID, err := strconv.Atoi(c.Params("alumni_id"))  
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    }

    p, err := repository.CountWorkAlumni(ctx, alumniID)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
    }

    return c.JSON(fiber.Map{"success": true, "data": p})
}



func GetPekerjaanByAlumniID(c *fiber.Ctx) error {
    ctx := c.Context()
    alumniID, err := strconv.Atoi(c.Params("alumni_id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Alumni ID tidak valid"})
    }

    list, err := repository.GetPekerjaanByAlumniID(ctx, alumniID)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pekerjaan"})
    }

    return c.JSON(fiber.Map{"success": true, "data": list})
}


func CreatePekerjaan(c *fiber.Ctx) error {
    ctx := c.Context()
    var req model.CreatePekerjaanRequest
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


// func DeletePekerjaan(c *fiber.Ctx) error {
//     ctx := c.Context()
//     id, err := strconv.Atoi(c.Params("id"))
//     if err != nil {
//         return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
//     }

//     rows, err := repository.DeletePekerjaan(ctx, id)
//     if err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Gagal hapus pekerjaan"})
//     }
//     if rows == 0 {
//         return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
//     }

//     return c.JSON(fiber.Map{"success": true, "message": "Pekerjaan berhasil dihapus"})
// }


func GetAllPekerjaanTrash(c *fiber.Ctx) error {
    ctx := c.Context()
    list, err := repository.GetAllPekerjaanTrash(ctx)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data pekerjaan"})
    }
    return c.JSON(fiber.Map{"success": true, "data": list})
}



func DeletePekerjaanTrash(c *fiber.Ctx) error {
    ctx := c.Context()
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    }

    rows, err := repository.DeletePekerjaanTrash(ctx, id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal hapus pekerjaan"})
    }
    if rows == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
    }

    return c.JSON(fiber.Map{"success": true, "message": "Pekerjaan berhasil dihapus"})
}




// func DeletePekerjaanTrashRestore(c *fiber.Ctx) error {
//     ctx := c.Context()
//     id, err := strconv.Atoi(c.Params("id"))
//     if err != nil {
//         return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
//     }

//     rows, err := repository.DeletePekerjaanTrashRestore(ctx, id)
//     if err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Gagal restore data trash pekerjaan"})
//     }
//     if rows == 0 {
//         return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan"})
//     }

//     return c.JSON(fiber.Map{"success": true, "message": "Pekerjaan berhasil di restore"})
// }






func RestorePekerjaanByUser(c *fiber.Ctx) error {
	pekerjaanID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "id pekerjaan tidak valid"})
	}

	role := c.Locals("role").(string)
	userID := c.Locals("user_id").(int)
	ctx := c.Context()
	var rowsAffected int64

	if role == "admin" {
		rowsAffected, err = repository.DeletePekerjaanTrashRestore(ctx, pekerjaanID, userID)
	} else {
		rowsAffected, err = repository.RestoreTrashPekerjaanByUser(ctx, pekerjaanID, userID)
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal restore pekerjaan"})
	}

	
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan atau Anda tidak memiliki akses"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Pekerjaan berhasil di restore",
	})
}


func HardDeletePekerjaanByUserInTrash(c *fiber.Ctx) error {
	pekerjaanID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "id pekerjaan tidak valid"})
	}

	role := c.Locals("role").(string)
	userID := c.Locals("user_id").(int)
	ctx := c.Context()
	var rowsAffected int64

	if role == "admin" {
		rowsAffected, err = repository.DeletePekerjaanTrash(ctx, pekerjaanID)
	} else {
		rowsAffected, err = repository.DeletePekerjaanTrashByUser(ctx, pekerjaanID, userID)
	}

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Gagal hard delete pekerjaan"})
	}

	
	if rowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "Pekerjaan tidak ditemukan atau Anda tidak memiliki akses"})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Pekerjaan berhasil di hard delete",
	})
}

