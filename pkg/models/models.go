package models

import (
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID           uint      `gorm:"primaryKey"`
	SessionToken string    `json:"session_token" gorm:"uuid;not null"`
	UserId       uint      `json:"user_email" gorm:"int;not null"`
	Expiration   time.Time `json:"expiration"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
type User struct {
	ID            uint   `gorm:"primaryKey"`
	GoogleID      string `json:"google_id" gorm:"varchar(30);not null"`
	UserEmail     string `json:"user_email" gorm:"int;not null"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type Task struct {
	ID         uint    `gorm:"primaryKey"`
	UserID     uint    `json:"user_id"` // Foreign key for User
	TaskStatus string  `json:"task_status"`
	Task       *string `json:"task,omitempty"`
	DropdownID *uint   `json:"category_id,omitempty"` // Foreign key for Category
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type JobApplication struct {
	ID                uint           `gorm:"primaryKey"`
	UserID            uint           `json:"user_id"` //Foreign key to User
	Company           string         `json:"company"`
	Title             string         `json:"title"`
	Location          *string        `json:"location,omitempty"`           // Pointer to support omission if empty
	ApplicationStatus *uint          `json:"application_status,omitempty"` // FK to dropdowns
	ApplicationType   *uint          `json:"application_type,omitempty"`   // FK to dropdown
	ResumeID          *uint          `json:"resume_id,omitempty"`          // Pointer to support omission if empty
	CoverLetterID     *uint          `json:"cover_letter_id,omitempty"`    // Pointer to support omission if empty
	PostingUrl        *string        `json:"posting_url,omitempty"`
	SalaryRange       *string        `json:"salary_range,omitempty"`
	ContactID         *uint          `json:"contact_id,omitempty"` // Pointer to support omission if empty
	Notes             *string        `json:"notes,omitempty"`
	CreatedAt         time.Time      `json:"created_at,omitempty"`
	UpdatedAt         time.Time      `json:"updated_at,omitempty"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type Document struct {
	ID               uint   `gorm:"primaryKey"`
	UserID           uint   // Foreign key for User
	OriginalFileName string `json:"original_file_name"`
	Path             string `json:"path"`
	TypeOfDocument   string `json:"type_of_document"`
	Notes            string `json:"notes"`
	DocumentName     string `json:"document_name"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

type Dropdown struct {
	ID        uint  `gorm:"primaryKey"`
	UserID    *uint `json:"user_id,omitempty"` // Foreign key for User
	Text      string
	TableType string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
