package models

import (
	"time"

	"gorm.io/gorm"
)

type Document struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint // Foreign key for User
	// User             *User `gorm:"foreignKey:UserID;references:ID"`
	OriginalFileName string
	Path             string
	TypeOfDocument   string
	Notes            string
	DocumentName     string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
