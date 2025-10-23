package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() (*mongo.Database, error){
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017" // Default jika .env tidak ada
		log.Println("Peringatan: MONGODB_URI tidak disetel. Menggunakan default:", mongoURI)
	}

	
	dbName := os.Getenv("DATABASE_NAME")
	if dbName == "" {
		log.Fatal("❌ DATABASE_NAME tidak ditemukan di .env")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("koneksi ke MongoDB gagal: %v", err)
	}

	err = client.Ping(ctx,nil)
	if err !=nil{
		return nil, fmt.Errorf("gagal konek ke MongoDB: %v", err)
	}

	fmt.Println("berhasil konek ke MongoDB")
	
	return client.Database(dbName), nil


	
}
