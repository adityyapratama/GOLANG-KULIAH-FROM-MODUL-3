package route

import (
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App,authSvc service.IAuthService, alumniSvc service.IAlumniService,pekerjaanSvc service.IPekerjaanService,) {
	RegisterAuthRoutes(app, authSvc)

	api := app.Group("api",middleware.AuthRequired())
	RegisterAlumniRoutes(api, alumniSvc)
	RegisterPekerjaanRoutes(api, pekerjaanSvc)


}
