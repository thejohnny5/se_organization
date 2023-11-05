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
	ID         uint    `gorm:"primaryKey"`
	UserID     uint    `json:"user_id"` // Foreign key for User
	TaskStatus string  `json:"task_status"`
	Task       *string `json:"task,omitempty"`
	CategoryID *uint   `json:"category_id,omitempty"` // Foreign key for Category
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type JobApplication struct {
	ID                uint           `gorm:"primaryKey"`
	UserID            uint           `json:"user_id"` //Foreign key to User
	Company           string         `json:"company"`
	Title             string         `json:"title"`
	Location          *string        `json:"location,omitempty"` // Pointer to support omission if empty
	ApplicationStatus *string        `json:"applicationStatus,omitempty"`
	ApplicationType   *string        `json:"applicationType,omitempty"`
	ResumeID          *uint          `json:"resumeId,omitempty"`      // Pointer to support omission if empty
	CoverLetterID     *uint          `json:"coverLetterId,omitempty"` // Pointer to support omission if empty
	PostingID         *string        `json:"postingId,omitempty"`
	SalaryRange       *string        `json:"salaryRange,omitempty"`
	ContactID         *uint          `json:"contactId,omitempty"` // Pointer to support omission if empty
	Notes             *string        `json:"notes,omitempty"`
	CreatedAt         time.Time      `json:"created_at,omitempty"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty"`
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
