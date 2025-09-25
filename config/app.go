package config

import (
	
	"golang-kuliah-from-modul-3/middleware"
	"golang-kuliah-from-modul-3/route"

	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
    app := fiber.New()
    
    app.Use(middleware.LoggerMiddleware)
    route.RegisterRoutes(app)

    return app
}
