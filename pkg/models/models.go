package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    string `json:"user_id" gorm:"varchar(30);not null"`
	UserEmail string `json:"user_email" gorm:"int;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Task struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `json:"user_id,omitempty"` // Foreign key for User
	TaskStatus string `json:"task_status"`
	Task       string `json:"task,omitempty"`
	CategoryID *uint  `json:"category_id,omitempty"` // Foreign key for Category
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type JobApplication struct {
	ID                uint `gorm:"primaryKey"`
	Company           string
	Title             string
	Location          string
	ApplicationStatus string
	ApplicationType   string
	ResumeID          uint // Foreign key for Document
	CoverLetterID     uint // Foreign key for Document
	PostingID         string
	SalaryRange       string
	ContactID         uint // This could be another model, representing the contact person
	Notes             string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type Document struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint // Foreign key for User
	UUID         int  `gorm:"uniqueIndex;not null"`
	Path         string
	Notes        string
	DocumentName string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Category struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint // Foreign key for User
	Category  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
