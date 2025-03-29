package workers

import (
	"database/sql"
	"fmt"
	"time"
)

func CleanupExpiredFiles(db *sql.DB) {
	for {
		_, err := db.Exec("DELETE FROM files WHERE upload_at < NOW() - INTERVAL '7 days'")
		if err != nil {
			fmt.Println("Failed to delete expired files:", err)
		} else {
			fmt.Println("Expired files deleted")
		}
		time.Sleep(1 * time.Hour)
	}
}
