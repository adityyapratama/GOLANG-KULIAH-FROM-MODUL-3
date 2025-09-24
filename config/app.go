package config

import (
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/middleware"
	"golang-kuliah-from-modul-3/route"

	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
    app := fiber.New()

    // Middleware global
    app.Use(middleware.LoggerMiddleware)

    // Route khusus check alumni
    app.Post("/check/:key", service.GetAllAlumni)
    
    route.RegisterRoutes(app)

    return app
}
