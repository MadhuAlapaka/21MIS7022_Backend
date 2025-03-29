package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"your_project/config"
	"your_project/internal/routes"
	"your_project/internal/utils"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database and Redis
	config.ConnectDB()
	config.ConnectRedis()

	// Initialize AWS S3 client
	utils.InitS3()

	// Setup router
	router := routes.SetupRoutes()

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
