package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    string `gorm:"uniqueIndex;not null"`
	UserEmail string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Task struct {
	ID         uint `gorm:"primaryKey"`
	UserID     uint // Foreign key for User
	TaskStatus string
	Task       string
	CategoryID uint // Foreign key for Category
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
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   // Foreign key for User
	UUID         string `gorm:"uniqueIndex;not null"`
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
