package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint
	User       User `gorm:"foreignKey:UserId;references:ID"`
	TaskStatus string
	Task       *string
	DropdownID *uint `json:"category_id,omitempty"` // Foreign key for Category
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
