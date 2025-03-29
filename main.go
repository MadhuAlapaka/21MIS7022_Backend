package main

import (
	"fmt"
	"log"
	"net/http"

	"your_project/config"
	"your_project/internal/routes"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitDB()
	config.InitRedis()

	router := routes.SetupRoutes()
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
