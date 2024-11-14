package services

import (
	"cv-form/models"
	"cv-form/repository"
	"cv-form/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type ProfileService interface {
	GetProfile(profileCode uint) (models.Profile, error)
	CreateProfile(profile utils.PayloadProfile) (models.Profile, error)
	UpdateProfile(profileCode uint, profile utils.PayloadProfile) (models.Profile, error)
	UpdateImageProfile(profileCode uint, photoUrl string) (string, error)
	GetImageProfile(profileCode uint) (string, error)
}

type profileService struct {
	profileRepository repository.ProfileRepository
}

func NewProfileService(profileRepository repository.ProfileRepository) ProfileService {
	return &profileService{profileRepository: profileRepository}
}

func (s *profileService) GetProfile(profileCode uint) (models.Profile, error) {
	return s.profileRepository.FindProfile(profileCode)
}

func (s *profileService) CreateProfile(payload utils.PayloadProfile) (models.Profile, error) {
	rand.Seed(time.Now().UnixNano())
	randomNumber := fmt.Sprintf("%08d", rand.Intn(100000000))
	number, _ := strconv.ParseUint(randomNumber, 10, 32)
	profile := models.Profile{
		ProfileCode:    uint(number),
		WantedJobTitle: payload.WantedJobTitle,
		FirstName:      payload.FirstName,
		LastName:       payload.LastName,
		Email:          payload.Email,
		Phone:          payload.Phone,
		Country:        payload.Country,
		City:           payload.City,
		Address:        payload.Address,
		PostalCode:     payload.PostalCode,
		DrivingLicense: payload.DrivingLicense,
		Nationality:    payload.Nationality,
		PlaceOfBirth:   payload.PlaceOfBirth,
		DateOfBirth:    payload.DateOfBirth,
	}
	return s.profileRepository.CreateProfile(profile)
}

func (s *profileService) UpdateProfile(ProfileCode uint, profile utils.PayloadProfile) (models.Profile, error) {
	foundProfile, err := s.profileRepository.FindProfile(ProfileCode)
	if err != nil {
		return foundProfile, errors.New("Profile not found")
	}

	foundProfile.WantedJobTitle = profile.WantedJobTitle
	foundProfile.FirstName = profile.FirstName
	foundProfile.LastName = profile.LastName
	foundProfile.Email = profile.Email
	foundProfile.Phone = profile.Phone
	foundProfile.Country = profile.Country
	foundProfile.City = profile.City
	foundProfile.Address = profile.Address
	foundProfile.PostalCode = profile.PostalCode
	foundProfile.DrivingLicense = profile.DrivingLicense
	foundProfile.Nationality = profile.Nationality
	foundProfile.PlaceOfBirth = profile.PlaceOfBirth
	foundProfile.DateOfBirth = profile.DateOfBirth

	updatedProfile, err := s.profileRepository.UpdateProfile(foundProfile)
	if err != nil {
		return updatedProfile, errors.New("Failed to update profile")
	}
	return updatedProfile, nil
}

func (s *profileService) UpdateImageProfile(profileCode uint, photoUrl string) (string, error) {
	imageData := photoUrl
	if strings.Contains(imageData, "base64,") {
		parts := strings.Split(imageData, ",")
		if len(parts) > 1 {
			imageData = parts[1]
		}
	}
	decodedImage, err := base64.StdEncoding.DecodeString(imageData)
	if err != nil {
		return "", errors.New("invalid base64 data")
	}

	fileName := fmt.Sprintf("./uploads/photo/%v.png", profileCode)
	err = ioutil.WriteFile(fileName, decodedImage, 0644)
	if err != nil {
		return "", errors.New("Failed to write Image profile")
	}

	_, err = s.profileRepository.UpdateImageProfile(profileCode, fileName)
	if err != nil {
		return "", errors.New("Failed to update Image profile")
	}
	return fileName, nil
}

func (s *profileService) GetImageProfile(profileCode uint) (string, error) {
	data, err := s.profileRepository.FindProfile(profileCode)
	if err != nil {
		return "", errors.New("Failed to find profile")
	}

	if data.PhotoUrl == "" {
		return "", errors.New("No Image profile found")
	}

	return data.PhotoUrl, nil
}
