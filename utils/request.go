package utils

type PayloadProfile struct {
	WantedJobTitle string `json:"wantedJobTitle" validate:"required,max=50"`
	FirstName      string `json:"firstName" validate:"required,max=20"`
	LastName       string `json:"lastName" validate:"required,max=20"`
	Email          string `json:"email" validate:"required,min=5,max=50,email"`
	Phone          string `json:"phone" validate:"required,min=9,max=20"`
	Country        string `json:"country" validate:"required,max=30"`
	City           string `json:"city" validate:"required,max=30"`
	Address        string `json:"address" validate:"required"`
	PostalCode     uint   `json:"postalCode" validate:"required"`
	DrivingLicense string `json:"drivingLicense" validate:"required,max=20"`
	Nationality    string `json:"nationality" validate:"required,max=30"`
	PlaceOfBirth   string `json:"placeOfBirth" validate:"required,max=30"`
	DateOfBirth    string `json:"dateOfBirth" validate:"required"`
	PhotoUrl       string `json:"photoUrl"`
}

type ParamsProfileCode struct {
	ProfileCode uint `json:"profileCode"`
}

type PayloadEmployment struct {
	JobTitle    string `json:"jobTitle" validate:"required,max=20"`
	Employer    string `json:"employer" validate:"required,max=20"`
	StartDate   string `json:"startDate" validate:"required,max=20"`
	EndDate     string `json:"endDate" validate:"required,max=20"`
	City        string `json:"city" validate:"required,max=20"`
	Description string `json:"description" validate:"required,max=100"`
}

type PhotoProfile struct {
	Base64img string `json:"base64img"`
}
