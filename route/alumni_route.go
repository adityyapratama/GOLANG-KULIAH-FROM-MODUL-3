package route

import (
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAlumniRoutes(router fiber.Router) {
    alumni := router.Group("/alumni")

    // alumni.Get("/", service.GetAllAlumni)
    alumni.Get("/:id", service.GetAlumniByID)

    alumni.Get("/", service.GetAllAlumniShorting)

    alumni.Post("/", middleware.AdminOnly(), service.CreateAlumni)
    alumni.Put("/:id", middleware.AdminOnly(), service.UpdateAlumni)
    alumni.Delete("/:id",middleware.AdminOnly(),service.DeleteAlumni)
}
