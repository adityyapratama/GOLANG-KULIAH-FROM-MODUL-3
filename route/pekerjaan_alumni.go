package route

import (
    "golang-kuliah-from-modul-3/app/service"
    "golang-kuliah-from-modul-3/middleware"

    "github.com/gofiber/fiber/v2"
)

func RegisterAlumniPekerjaanRoutes(router fiber.Router) {
    alumni := router.Group("pekerjaan")

    
    alumni.Get("/", service.GetAllPekerjaan)
    alumni.Post("/", middleware.AdminOnly(), service.CreatePekerjaan)
    
    
    alumni.Get("/trash", service.GetAllPekerjaanTrash)
    alumni.Delete("/trash/:id", service.DeletePekerjaanTrash)
    alumni.Put("/trash/restore/:id", service.RestorePekerjaanByUser)
    alumni.Delete("/trash/:id",service.HardDeletePekerjaanByUserInTrash)
    
    alumni.Get("/total/kerja/:alumni_id", service.GetTotalKerja)
    alumni.Get("/alumni/:alumni_id", middleware.AdminOnly(), service.GetPekerjaanByAlumniID)
    


    alumni.Get("/:id", service.GetPekerjaanByID)
    alumni.Put("/:id", middleware.AdminOnly(), service.UpdatePekerjaan)
    alumni.Delete("/:id", service.DeletePekerjaanByUser)



}