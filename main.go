package main

import (
	"belajar/golang/database"
	"belajar/golang/router"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	port := os.Getenv("PORT")
	r.Run(": " + port)
}
