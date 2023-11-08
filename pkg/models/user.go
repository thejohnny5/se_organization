package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint   `gorm:"primaryKey"`
	GoogleID      string `gorm:"varchar(30);not null"`
	UserEmail     string
	VerifiedEmail bool
	Picture       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
