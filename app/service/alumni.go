package service

import (
	"golang-kuliah-from-modul-3/app/model"
	"golang-kuliah-from-modul-3/app/repository"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)


func GetAllAlumni(c *fiber.Ctx) error {
    ctx := c.Context()
    list, err := repository.GetAllAlumni(ctx)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data alumni"})
    }
    return c.JSON(fiber.Map{"success": true, "data": list})
}


func GetAlumniByID(c *fiber.Ctx) error {
    ctx := c.Context()
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    }

    alumni, err := repository.GetAlumniByID(ctx,id)
    if err != nil {
        return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
    }

    return c.JSON(fiber.Map{"success": true, "data": alumni})
}


func CreateAlumni(c *fiber.Ctx) error {
    ctx := c.Context()
    var req model.Alumni
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
    }

    if req.NIM == "" || req.Nama == "" || req.Email == "" {
        return c.Status(400).JSON(fiber.Map{"error": "Field wajib diisi"})
    }

    if err := repository.CreateAlumni(ctx,&req); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal menambah alumni"})
    }

    req.CreatedAt = time.Now()
    req.UpdatedAt = time.Now()

    return c.Status(201).JSON(fiber.Map{"success": true, "data": req})
}


func UpdateAlumni(c *fiber.Ctx) error {
    ctx := c.Context()
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    }

    var req model.Alumni
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
    }
    req.ID = id

    rows, err := repository.UpdateAlumni(ctx,&req)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal update alumni"})
    }
    if rows == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
    }

    req.UpdatedAt = time.Now()

    return c.JSON(fiber.Map{"success": true, "data": req})
}


func DeleteAlumni(c *fiber.Ctx) error {
    ctx := c.Context()
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
    }

    rows, err := repository.DeleteAlumni(ctx, id)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Gagal hapus alumni"})
    }
    if rows == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "Alumni tidak ditemukan"})
    }

    return c.JSON(fiber.Map{"success": true, "message": "Alumni berhasil dihapus"})
}
