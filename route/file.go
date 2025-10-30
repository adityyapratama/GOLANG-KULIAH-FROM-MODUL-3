package route

import (
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterFileUpload(router fiber.Router, FileSvc service.FileService) {
	
	file := router.Group("/files")
	file.Post("/upload", FileSvc.UploadFile) 
	file.Get("/", FileSvc.GetAllFiles)       
	file.Get("/:id", FileSvc.GetFileByID)    
	file.Delete("/:id", FileSvc.DeleteFile)  

	// untuk user dan admin
	file.Post("/foto", FileSvc.UploadFoto)             
	file.Post("/sertifikat", FileSvc.UploadSertifikat) 

	
	
	admin := router.Group("/admin")
	admin.Use(middleware.AdminOnly()) 

	admin.Post("/users/:user_id/foto", FileSvc.UploadFoto)             
	admin.Post("/users/:user_id/sertifikat", FileSvc.UploadSertifikat) 
}
