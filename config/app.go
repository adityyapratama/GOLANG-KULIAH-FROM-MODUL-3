package config

import (
    "golang-kuliah-from-modul-3/app/service" // <-- Import service
    "golang-kuliah-from-modul-3/middleware"
    "golang-kuliah-from-modul-3/route"

    "github.com/gofiber/fiber/v2"
)


func NewApp(authSvc service.IAuthService,
			 alumniSvc service.IAlumniService,
			 pekerjaanSvc service.IPekerjaanService,			
			 ) *fiber.App {


    app := fiber.New()
    app.Use(middleware.LoggerMiddleware)

    
    route.RegisterRoutes(app, authSvc, alumniSvc, pekerjaanSvc)
    

    return app
}