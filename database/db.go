package database

import (
	"belajar/golang/model"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func StartDB() {
	err := godotenv.Load() // Load environment variables from .env
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	fmt.Println("Sukses koneksi ke database")
	db.Debug().AutoMigrate(model.User{}, model.Photo{}, model.Comment{}, model.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
