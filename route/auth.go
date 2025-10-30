package route

import (
	"golang-kuliah-from-modul-3/app/service"
	

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(app *fiber.App, authSvc service.IAuthService) {

	
	app.Post("/login", authSvc.Login)
	app.Post("/register", authSvc.Register)
	
}

