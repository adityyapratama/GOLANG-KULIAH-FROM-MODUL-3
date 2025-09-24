package route

import (
	"golang-kuliah-from-modul-3/app/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterAlumniPekerjaanRoutes(app *fiber.App) {
    alumni := app.Group("pekerjaan")

    alumni.Get("/", service.GetAllPekerjaan)
    alumni.Get("/:id", service.GetPekerjaanByID)

    alumni.Get("/alumni/:alumni_id", service.GetPekerjaanByAlumniID)
    alumni.Post("/", service.CreatePekerjaan)
    alumni.Put("/:id", service.UpdatePekerjaan)
    alumni.Delete("/:id", service.DeletePekerjaan)
}
