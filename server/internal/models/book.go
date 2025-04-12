package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        int64          `gorm:"primaryKey"`
	Title     string         `gorm:"size:255;not null"`
	Author    string         `gorm:"size:255;not null"`
	ISBN      string         `gorm:"size:20;not null;uniqueIndex"`
	Price     float64        `gorm:"not null"`
	Stock     int32          `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&Book{})
}
