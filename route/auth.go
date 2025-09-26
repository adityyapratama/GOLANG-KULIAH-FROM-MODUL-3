package route

import (
	"golang-kuliah-from-modul-3/app/service"
	

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App) {
	// Rute publik untuk login
	app.Post("/login", service.Login)

	// Rute publik untuk registrasi user baru
	app.Post("/register", service.Register)
	app.Get("/register", service.GetProfile)
	app.Get("/users", service.GetUsersService)

	
}

