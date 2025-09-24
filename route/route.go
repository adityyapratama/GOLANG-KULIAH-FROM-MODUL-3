package route

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	RegisterAlumniRoutes(app)
	RegisterMahasiswaRoutes(app)
	RegisterAlumniPekerjaanRoutes(app)
}
