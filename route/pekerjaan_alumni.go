package route

import (
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAlumniPekerjaanRoutes(router fiber.Router) {
    alumni := router.Group("pekerjaan")

    alumni.Get("/", service.GetAllPekerjaan)
    alumni.Get("/:id", service.GetPekerjaanByID)

    alumni.Get("/alumni/:alumni_id", middleware.AdminOnly(), service.GetPekerjaanByAlumniID)
    alumni.Post("/",  middleware.AdminOnly(),service.CreatePekerjaan)
    alumni.Put("/:id", middleware.AdminOnly(),service.UpdatePekerjaan)
    alumni.Delete("/:id",service.DeletePekerjaanByUser)


    
    

    
    
}
