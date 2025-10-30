package route

import (
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/middleware"

	"github.com/gofiber/fiber/v2"
)



func RegisterAlumniRoutes(router fiber.Router, alumniSvc service.IAlumniService) {
	alumni := router.Group("/alumni")
	alumni.Post("/", middleware.AdminOnly(), alumniSvc.CreateAlumni)
    alumni.Get("/", alumniSvc.GetAllAlumni) 
    alumni.Get("/:id", alumniSvc.GetAlumniByID) 
    alumni.Put("/:id", middleware.AdminOnly(), alumniSvc.UpdateAlumni)
    alumni.Delete("/:id", middleware.AdminOnly(), alumniSvc.DeleteAlumni)


}