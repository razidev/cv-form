package models

import "time"

type Profile struct {
	ID             uint   `gorm:"primary_key"`
	ProfileCode    uint   `gorm:"unique"`
	WantedJobTitle string `gorm:"not null"`
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	Email          string `gorm:"unique;not null"`
	Phone          string `gorm:"not null"`
	Country        string `gorm:"not null"`
	City           string `gorm:"not null"`
	Address        string `gorm:"not null"`
	PostalCode     uint   `gorm:"not null"`
	DrivingLicense string `gorm:"not null"`
	Nationality    string `gorm:"not null"`
	PlaceOfBirth   string `gorm:"not null"`
	DateOfBirth    string `gorm:"not null"`
	PhotoUrl       string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
