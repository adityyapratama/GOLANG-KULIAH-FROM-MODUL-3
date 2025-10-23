package route

import (
	"golang-kuliah-from-modul-3/app/service"
	

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App, authSvc service.IAuthService) {
	// Rute publik untuk login
	app.Post("/login", authSvc.Login)
	app.Post("/register", authSvc.Register)



	// app.Get("/register", service.GetProfile)
	// app.Get("/users", service.GetUsersService)

	
}

