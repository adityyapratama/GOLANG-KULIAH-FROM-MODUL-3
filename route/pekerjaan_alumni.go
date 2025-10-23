package route

import (
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/middleware"

	"github.com/gofiber/fiber/v2"
)


func RegisterPekerjaanRoutes(router fiber.Router, pekerjaanSvc service.IPekerjaanService) {
    

    pekerjaan := router.Group("/pekerjaan")


    pekerjaan.Get("/", pekerjaanSvc.GetAllPekerjaan)
    pekerjaan.Get("/:id", pekerjaanSvc.GetPekerjaanByID)
    pekerjaan.Get("/alumni/:alumni_id", middleware.AdminOnly(), pekerjaanSvc.GetPekerjaanByAlumniID)
    pekerjaan.Post("/", middleware.AdminOnly(), pekerjaanSvc.CreatePekerjaan)
    pekerjaan.Put("/:id", middleware.AdminOnly(), pekerjaanSvc.UpdatePekerjaan)
    pekerjaan.Delete("/:id", middleware.AdminOnly(), pekerjaanSvc.DeletePekerjaan)
}