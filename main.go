package main

import (
	"golang-kuliah-from-modul-3/app/repository"
	"golang-kuliah-from-modul-3/app/service"
	"golang-kuliah-from-modul-3/config"
	"golang-kuliah-from-modul-3/database"

	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title           Mahasiswa, Alumni, pekerjaan_alumni and upload file alumni API Documentation
// @version         1.0
// @description     Dokumentasi API untuk Mahasiswa
// @host      localhost:3000
// @BasePath  /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @schemes http
func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Tidak menemukan file .env, pakai environment system")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("❌ Gagal konek DB:", err)
	}

	authRepo := repository.NewAuthRepository(db)
	alumniRepo := repository.NewAlumniRepository(db)
	pekerjaanRepo := repository.NewPekerjaanRepository(db)
	FileRepo := repository.NewRepositoryFile(db)

	authSvc := service.NewAuthService(authRepo)
	alumniSvc := service.NewAlumniService(alumniRepo)
	pekerjaanSvc := service.NewPekerjaanService(pekerjaanRepo, alumniRepo)

	// Set upload path untuk file service
	uploadPath := os.Getenv("UPLOAD_PATH")
	if uploadPath == "" {
		uploadPath = "./uploads" // Default path
	}
	FileSvc := service.NewFileService(FileRepo, uploadPath)

	app := config.NewApp(authSvc, alumniSvc, pekerjaanSvc, FileSvc)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
