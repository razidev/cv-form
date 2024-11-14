package models

import "time"

type Employment struct {
	ID          uint   `gorm:"primary_key"`
	ProfileCode uint   `gorm:"not null"`
	JobTitle    string `gorm:"not null"`
	Employer    string `gorm:"not null"`
	StartDate   string `gorm:"not null"`
	EndDate     string `gorm:"not null"`
	City        string `gorm:"not null"`
	Description string `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
