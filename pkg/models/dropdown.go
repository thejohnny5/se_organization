package models

import (
	"time"

	"gorm.io/gorm"
)

type Dropdown struct {
	ID     uint `gorm:"primaryKey"`
	UserId uint
	// User      User `gorm:"foreignKey:UserId;references:ID"` // Foreign key for User
	Text      string
	TableType string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
