package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB()  {
	dsn := "host=localhost user=postgres password=password dbname=BackendGolangg port=5433 sslmode=disable"
	
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("gagal koneksi ke database",err)
	}

	if err = DB.Ping(); err != nil {
	log.Fatal("Gagal ping database:", err)
}
	fmt.Println("Berhasil terhubung ke database PostgreSQL")
}
