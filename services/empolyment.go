package services

import (
	"cv-form/models"
	"cv-form/repository"
	"cv-form/utils"
	"errors"
)

type EmploymentService interface {
	CreateEmployment(profileCode uint, employee utils.PayloadEmployment) (models.Employment, error)
	ListEmployments(profileCode uint) ([]models.Employment, error)
	DeleteEmployment(id uint, profileCode uint) error
}

type employmentService struct {
	employmentRepository repository.EmploymentRepository
	profileRepository    repository.ProfileRepository
}

func NewEmploymentService(employmentRepository repository.EmploymentRepository, profileRepository repository.ProfileRepository) EmploymentService {
	return &employmentService{
		employmentRepository: employmentRepository,
		profileRepository:    profileRepository,
	}
}

func (s *employmentService) CreateEmployment(profileCode uint, payload utils.PayloadEmployment) (models.Employment, error) {
	employment := models.Employment{
		ProfileCode: profileCode,
		JobTitle:    payload.JobTitle,
		Employer:    payload.Employer,
		StartDate:   payload.StartDate,
		EndDate:     payload.EndDate,
		City:        payload.City,
		Description: payload.Description,
	}

	return s.employmentRepository.CreateEmployment(&employment)
}

func (s *employmentService) ListEmployments(profileCode uint) ([]models.Employment, error) {
	_, err := s.profileRepository.FindProfile(profileCode)
	if err != nil {
		return nil, errors.New("Profile not found")
	}

	employments, err := s.employmentRepository.GetEmployments(profileCode)
	if err != nil {
		return nil, err
	}

	return employments, nil
}

func (s *employmentService) DeleteEmployment(id uint, profileCode uint) error {
	_, err := s.profileRepository.FindProfile(profileCode)
	if err != nil {
		return errors.New("Profile not found")
	}
	_, err = s.employmentRepository.FindEmployments(id)
	if err != nil {
		return errors.New("Employment not found")
	}

	return s.employmentRepository.DeleteEmployment(id)
}
