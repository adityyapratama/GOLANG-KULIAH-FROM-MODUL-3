package main

import (
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

	
	if err := database.ConnectDB(); err != nil {
		log.Fatal("❌ Gagal konek DB:", err)
	}
	defer database.DB.Close()

	
	app := config.NewApp()
	

	
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
