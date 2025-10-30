package config

import (
	"golang-kuliah-from-modul-3/app/service" // <-- Import service
	"golang-kuliah-from-modul-3/middleware"
	"golang-kuliah-from-modul-3/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewApp(authSvc service.IAuthService,
	alumniSvc service.IAlumniService,
	pekerjaanSvc service.IPekerjaanService,
	FileSvc service.FileService,
) *fiber.App {

	
	// 10MB untuk limit global (bisa di-override di handler spesifik)
	app := fiber.New(fiber.Config{
		BodyLimit: 10 * 1024 * 1024, // 10MB - limit maksimal untuk semua request
	})

	// Middleware CORS untuk akses dari frontend
	app.Use(cors.New())

	// Middleware Logger custom
	app.Use(middleware.LoggerMiddleware)

	// Serve static files untuk download/akses uploaded files
	app.Static("/uploads", "./uploads")

	// Register semua routes
	route.RegisterRoutes(app, authSvc, alumniSvc, pekerjaanSvc, FileSvc)

	return app
}
