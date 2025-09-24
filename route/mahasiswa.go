package route

import (
	"golang-kuliah-from-modul-3/app/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterMahasiswaRoutes(app *fiber.App) {
    alumni := app.Group("/mahasiswa")

    alumni.Get("/", service.GetAllMahasiswa)
    alumni.Get("/:id", service.GetMahasiswaByID)
    alumni.Post("/", service.CreateMahasiswa)
    alumni.Put("/:id", service.UpdateMahasiswa)
    alumni.Delete("/:id", service.DeleteMahasiswa)
}
