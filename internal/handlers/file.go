package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"your_project/config"
	"your_project/internal/models"
	"your_project/internal/utils"

	"github.com/gorilla/mux"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.Header.Get("User-ID"))

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Upload file to S3
	s3URL, err := utils.UploadToS3(file, header.Filename)
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}

	// Save metadata in DB
	fileMetadata := models.FileMetadata{
		UserID:     uint(userID),
		FileName:   header.Filename,
		Size:       header.Size,
		S3URL:      s3URL,
		UploadDate: time.Now().Format(time.RFC3339),
	}

	config.DB.Create(&fileMetadata)

	// Cache metadata in Redis
	cacheKey := fmt.Sprintf("file:%d", fileMetadata.ID)
	config.RedisClient.Set(r.Context(), cacheKey, fileMetadata, time.Minute*5)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "File uploaded successfully",
		"url":     s3URL,
	})
}

func GetFilesHandler(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.Atoi(r.Header.Get("User-ID"))

	var files []models.FileMetadata
	config.DB.Where("user_id = ?", userID).Find(&files)

	json.NewEncoder(w).Encode(files)
}

func ShareFileHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileID := params["file_id"]

	var file models.FileMetadata
	config.DB.First(&file, fileID)

	if file.ID == 0 {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"shared_url": file.S3URL,
	})
}
