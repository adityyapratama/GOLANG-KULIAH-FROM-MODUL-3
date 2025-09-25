package route

import (
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterMahasiswaRoutes(router fiber.Router) {
    alumni := router.Group("/mahasiswa")

    alumni.Get("/", service.GetAllMahasiswa)
    alumni.Get("/:id", service.GetMahasiswaByID)
    alumni.Post("/", middleware.AdminOnly(),service.CreateMahasiswa)
    alumni.Put("/:id", middleware.AdminOnly(),service.UpdateMahasiswa)
    alumni.Delete("/:id",middleware.AdminOnly(),service.DeleteMahasiswa)
}
