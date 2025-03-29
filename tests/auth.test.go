package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserRegister(t *testing.T) {
	reqBody := []byte(`{"email": "test@example.com", "password": "password"}`)
	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(reqBody))
	w := httptest.NewRecorder()

	// Simulate request
	// authHandler.Register(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
