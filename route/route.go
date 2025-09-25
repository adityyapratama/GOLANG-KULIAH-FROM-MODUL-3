package route

import (
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {

	RegisterAuthRoutes(app)
	api := app.Group("/api", middleware.AuthRequired())
	
	api.Get("/profile", service.GetProfile)

	RegisterAlumniRoutes(api)
	RegisterMahasiswaRoutes(api)
	RegisterAlumniPekerjaanRoutes(api)

	RegisterAuthRoutes(app)
}
