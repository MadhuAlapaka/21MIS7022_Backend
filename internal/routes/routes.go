package routes

import (
	"net/http"

	"your_project/internal/handlers"
	"your_project/internal/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Authentication routes
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Protected routes
	api := router.PathPrefix("/api").Subrouter()
	api.Use(middleware.AuthMiddleware)

	api.HandleFunc("/upload", handlers.UploadFileHandler).Methods("POST")
	api.HandleFunc("/files", handlers.GetFilesHandler).Methods("GET")
	api.HandleFunc("/share/{file_id}", handlers.ShareFileHandler).Methods("GET")

	return router
}
