package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type FileMetadata struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"index"`
	FileName   string `gorm:"not null"`
	Size       int64
	S3URL      string
	UploadDate string
}
