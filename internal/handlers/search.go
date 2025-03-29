package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type File struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Size     int64  `json:"size"`
	UploadAt string `json:"upload_at"`
	URL      string `json:"url"`
}

func SearchFiles(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := "SELECT id, name, type, size, upload_at, url FROM files WHERE 1=1"
		var args []interface{}

		name := r.URL.Query().Get("name")
		if name != "" {
			query += " AND name ILIKE '%' || $1 || '%'"
			args = append(args, name)
		}

		fileType := r.URL.Query().Get("type")
		if fileType != "" {
			query += " AND type = $" + string(len(args)+1)
			args = append(args, fileType)
		}

		rows, err := db.Query(query, args...)
		if err != nil {
			http.Error(w, "Failed to fetch files", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var files []File
		for rows.Next() {
			var f File
			if err := rows.Scan(&f.ID, &f.Name, &f.Type, &f.Size, &f.UploadAt, &f.URL); err != nil {
				http.Error(w, "Error scanning file", http.StatusInternalServerError)
				return
			}
			files = append(files, f)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	}
}
