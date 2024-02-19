package model

import (
	"time"

	"gorm.io/gorm"
)

type Asset struct {
	gorm.Model
	Name     string `gorm:"type:varchar(300);not null"`
	FileName string `gorm:"type:varchar(200);not null"`
	FileType string `gorm:"type:varchar(50);not null"`
	FilePath string `gorm:"type:varchar(100);not null"`
	FileSize string `gorm:"type:varchar(50);not null"`
}

type AssetRequest struct {
	Name     string
	FileName string
	FileType string
	FilePath string
	FileSize string
}
type AssetResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	FileName  string    `json:"fileName"`
	FileType  string    `json:"fileType"`
	FilePath  string    `json:"filePath"`
	FileSize  string    `json:"fileSize"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
