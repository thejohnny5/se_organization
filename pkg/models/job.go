package models

import (
	"time"

	"gorm.io/gorm"
)

type JobApplication struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint //Foreign key to User
	// User                User `gorm:"foreignKey:UserID;references:ID"`
	Company             string
	Title               string
	Location            *string // Pointer to support omission if empty
	ApplicationStatusId *uint
	ApplicationStatus   Dropdown `gorm:"foreignKey:ApplicationStatusId;references:ID"` // FK to dropdowns
	ApplicationTypeId   *uint
	ApplicationType     Dropdown `gorm:"foreignKey:ApplicationTypeId;references:ID"` // FK to dropdown
	ResumeId            *uint    // Pointer to support omission if empty
	Resume              Document `gorm:"foreignKey:ResumeId;references:ID"`
	CoverLetterId       *uint    // Pointer to support omission if empty
	CoverLetter         Document `gorm:"foreignKey:CoverLetterId;references:ID"`
	PostingUrl          *string
	SalaryLow           *int
	SalaryHigh          *int
	ContactID           *uint // Pointer to support omission if empty
	Notes               *string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}
