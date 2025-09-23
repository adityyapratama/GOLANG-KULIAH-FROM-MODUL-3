package route

import (
	"database/sql"
	"golang-kuliah-from-modul-3/app/service"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, db *sql.DB) {
	app.Post("/check/:key", func(c *fiber.Ctx) error {
		return service.CheckAlumniService(c, db)
	})
}
