package route

import (
	"golang-kuliah-from-modul-3/app/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterAlumniRoutes(app *fiber.App) {
    alumni := app.Group("/alumni")

    alumni.Get("/", service.GetAllAlumni)
    alumni.Get("/:id", service.GetAlumniByID)
    alumni.Post("/", service.CreateAlumni)
    alumni.Put("/:id", service.UpdateAlumni)
    alumni.Delete("/:id", service.DeleteAlumni)
}
