package utils

import "cv-form/models"

type Profile struct {
	ProfileCode    uint   `json:"profileCode"`
	WantedJobTitle string `json:"wantedJobTitle"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	Country        string `json:"country"`
	City           string `json:"city"`
	Address        string `json:"address"`
	PostalCode     uint   `json:"postalCode"`
	DrivingLicense string `json:"drivingLicense"`
	Nationality    string `json:"nationality"`
	PlaceOfBirth   string `json:"placeOfBirth"`
	DateOfBirth    string `json:"dateOfBirth"`
	PhotoUrl       string `json:"photoUrl"`
}

func ProfileResponse(profile models.Profile) *Profile {
	return &Profile{
		ProfileCode:    profile.ProfileCode,
		WantedJobTitle: profile.WantedJobTitle,
		FirstName:      profile.FirstName,
		LastName:       profile.LastName,
		Email:          profile.Email,
		Phone:          profile.Phone,
		Country:        profile.Country,
		City:           profile.City,
		Address:        profile.Address,
		PostalCode:     profile.PostalCode,
		DrivingLicense: profile.DrivingLicense,
		Nationality:    profile.Nationality,
		PlaceOfBirth:   profile.PlaceOfBirth,
		DateOfBirth:    profile.DateOfBirth,
		PhotoUrl:       profile.PhotoUrl,
	}
}

type Employment struct {
	ID          uint   `json:"id"`
	JobTitle    string `json:"jobTitle"`
	Employer    string `json:"employer"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	City        string `json:"city"`
	Description string `json:"description"`
}

func EmploymentResponse(employment models.Employment) Employment {
	return Employment{
		ID:          employment.ID,
		JobTitle:    employment.JobTitle,
		Employer:    employment.Employer,
		StartDate:   employment.StartDate,
		EndDate:     employment.EndDate,
		City:        employment.City,
		Description: employment.Description,
	}
}

func EmploymentResponses(employments []models.Employment) []Employment {
	var responses []Employment
	for _, employment := range employments {
		responses = append(responses, EmploymentResponse(employment))
	}

	return responses
}
