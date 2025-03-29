package main

import (
	"fmt"
	"time"

	"your_project/config"
)

func DeleteExpiredFiles() {
	for {
		time.Sleep(24 * time.Hour) // Run daily

		var expiredFiles []string
		config.DB.Raw("SELECT s3_url FROM file_metadata WHERE upload_date < NOW() - INTERVAL '30 days'").Scan(&expiredFiles)

		for _, url := range expiredFiles {
			// Delete from S3
			fmt.Println("Deleting file:", url)
			// Perform S3 delete operation here...
		}

		// Delete from DB
		config.DB.Exec("DELETE FROM file_metadata WHERE upload_date < NOW() - INTERVAL '30 days'")
	}
}
