package repository

import (
	"cv-form/config"
	"cv-form/models"
)

type ProfileRepository interface {
	CreateProfile(profile models.Profile) (models.Profile, error)
	FindProfile(ProfileCode uint) (models.Profile, error)
	UpdateProfile(profile models.Profile) (models.Profile, error)
	UpdateImageProfile(ProfileCode uint, photoUrl string) (models.Profile, error)
}

type profileRepository struct{}

func NewProfileRepository() ProfileRepository {
	return &profileRepository{}
}

func (r *profileRepository) CreateProfile(profile models.Profile) (models.Profile, error) {
	if err := config.DB.Create(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}

func (r *profileRepository) FindProfile(ProfileCode uint) (models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("profile_code = ?", ProfileCode).First(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}

func (r *profileRepository) UpdateProfile(profile models.Profile) (models.Profile, error) {
	if err := config.DB.Save(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}

func (r *profileRepository) UpdateImageProfile(ProfileCode uint, photoUrl string) (models.Profile, error) {
	var profile models.Profile
	if err := config.DB.Where("profile_code = ?", ProfileCode).First(&profile).Error; err != nil {
		return profile, err
	}

	profile.PhotoUrl = photoUrl
	if err := config.DB.Save(&profile).Error; err != nil {
		return profile, err
	}

	return profile, nil
}
