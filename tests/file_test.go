package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFileUpload(t *testing.T) {
	// Simulating file upload
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(map[string]string{
		"file_name": "testfile.txt",
		"content":   "This is a test file",
	})

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	// Simulate request to the file upload handler
	// fileHandler.UploadFile(w, req)  // Uncomment when integrating with handler

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var response map[string]string
	json.Unmarshal(w.Body.Bytes(), &response)

	if response["message"] != "File uploaded successfully" {
		t.Errorf("Unexpected response message: %s", response["message"])
	}
}

func TestFileRetrieval(t *testing.T) {
	req := httptest.NewRequest("GET", "/files", nil)
	w := httptest.NewRecorder()

	// Simulate request to the file retrieval handler
	// fileHandler.GetFiles(w, req)  // Uncomment when integrating with handler

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	var files []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &files)

	if len(files) == 0 {
		t.Errorf("Expected at least one file, got none")
	}
}
