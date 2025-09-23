package main

import (
	"golang-kuliah-from-modul-3/config"
	"golang-kuliah-from-modul-3/database"
	"golang-kuliah-from-modul-3/route"
	"os"
)

func main() {
	database.ConnectDB()
	defer database.DB.Close()

	
	app := config.NewApp()
	route.RegisterRoutes(app, db)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}
	app.Listen(":" + port)
}
