package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID           uint   `gorm:"primaryKey"`
	SessionToken string `gorm:"uuid;not null"`
	UserId       uint
	// User         User `gorm:"foreignKey:UserId;references:ID"`
	Expiration time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
