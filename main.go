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

func main() {
	
	
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ Tidak menemukan file .env, pakai environment system")
	}

	db, err :=database.ConnectDB()
	if err != nil {
		log.Fatal("❌ Gagal konek DB:", err)
	}

	authRepo := repository.NewAuthRepository(db)
	alumniRepo := repository.NewAlumniRepository(db)
	pekerjaanRepo := repository.NewPekerjaanRepository(db)

	authSvc := service.NewAuthService(authRepo)

	alumniSvc := service.NewAlumniService(alumniRepo)
	pekerjaanSvc := service.NewPekerjaanService(pekerjaanRepo, alumniRepo)

	app := config.NewApp(authSvc, alumniSvc,pekerjaanSvc)

	
	
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
