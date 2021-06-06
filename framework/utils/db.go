package utils

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/octavioduarte/simple-api-with-go/domain"
)

func ConnectDB() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error when loading .env file")
	}

	connection := os.Getenv("connectiondb")

	db, err := gorm.Open("postgres", connection)

	if err != nil {
		log.Fatalf("Error on connection DB: %v", err)
		// Kill application if connection DB throws
		panic(err)
	}

	db.AutoMigrate(&domain.User{})

	return db
}
