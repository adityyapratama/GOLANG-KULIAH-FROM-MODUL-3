package config

import (
	"database/sql"
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewApp(db *sql.DB) *fiber.App {
	app := fiber.New()
	app.Use(middleware.LoggerMiddleware)
	app.Post("/check/:key", func(c *fiber.Ctx) error {
		return service.(c, db)
	})
	return app
}
