package models

import (	
	"time"
)

type FileMetadata struct {
	ID         uint      `gorm:"primaryKey"`
	FileName   string    `gorm:"unique"`
	UploadTime time.Time
}