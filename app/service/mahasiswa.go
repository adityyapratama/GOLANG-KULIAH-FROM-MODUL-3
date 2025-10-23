package service

// import (
// 	"golang-kuliah-from-modul-3/app/model"
// 	"golang-kuliah-from-modul-3/app/repository"
// 	"strconv"
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// )


// func GetAllMahasiswa(c *fiber.Ctx) error {
//     ctx := c.Context()
//     list, err := repository.GetAllMahasiswa(ctx)
//     if err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Gagal mengambil data mahasiswa"})
//     }
//     return c.JSON(fiber.Map{"success": true, "data": list})
// }


// func GetMahasiswaByID(c *fiber.Ctx) error {
//     ctx := c.Context()
//     id, err := strconv.Atoi(c.Params("id"))
//     if err != nil {
//         return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
//     }

//     m, err := repository.GetMahasiswaByID(ctx, id)
//     if err != nil {
//         return c.Status(404).JSON(fiber.Map{"error": "Mahasiswa tidak ditemukan"})
//     }

//     return c.JSON(fiber.Map{"success": true, "data": m})
// }


// func CreateMahasiswa(c *fiber.Ctx) error {
//     ctx := c.Context()
//     var req model.Mahasiswa
//     if err := c.BodyParser(&req); err != nil {
//         return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
//     }

//     if req.NIM == "" || req.Nama == "" || req.Jurusan == "" || req.Email == "" {
//         return c.Status(400).JSON(fiber.Map{"error": "Semua field harus diisi"})
//     }

//     if err := repository.CreateMahasiswa(ctx, &req); err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Gagal menambah mahasiswa"})
//     }

//     req.CreatedAt = time.Now()
//     req.UpdatedAt = time.Now()

//     return c.Status(201).JSON(fiber.Map{"success": true, "data": req})
// }


// func UpdateMahasiswa(c *fiber.Ctx) error {
//     ctx := c.Context()
//     id, err := strconv.Atoi(c.Params("id"))
//     if err != nil {
//         return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
//     }

//     var req model.Mahasiswa
//     if err := c.BodyParser(&req); err != nil {
//         return c.Status(400).JSON(fiber.Map{"error": "Request body tidak valid"})
//     }
//     req.ID = id

//     rows, err := repository.UpdateMahasiswa(ctx, &req)
//     if err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Gagal update mahasiswa"})
//     }
//     if rows == 0 {
//         return c.Status(404).JSON(fiber.Map{"error": "Mahasiswa tidak ditemukan"})
//     }

//     req.UpdatedAt = time.Now()

//     return c.JSON(fiber.Map{"success": true, "data": req})
// }


// func DeleteMahasiswa(c *fiber.Ctx) error {
//     ctx := c.Context()
//     id, err := strconv.Atoi(c.Params("id"))
//     if err != nil {
//         return c.Status(400).JSON(fiber.Map{"error": "ID tidak valid"})
//     }

//     rows, err := repository.DeleteMahasiswa(ctx, id)
//     if err != nil {
//         return c.Status(500).JSON(fiber.Map{"error": "Gagal hapus mahasiswa"})
//     }
//     if rows == 0 {
//         return c.Status(404).JSON(fiber.Map{"error": "Mahasiswa tidak ditemukan"})
//     }

//     return c.JSON(fiber.Map{"success": true, "message": "Mahasiswa berhasil dihapus"})
// }
