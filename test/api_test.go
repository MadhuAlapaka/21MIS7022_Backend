package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"your_project/internal/routes"
)

func TestRegister(t *testing.T) {
	router := routes.SetupRoutes()

	payload := []byte(`{"email": "test@example.com", "password": "password123"}`)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, status)
	}
}

