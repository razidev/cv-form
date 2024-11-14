package repository

import (
	"cv-form/config"
	"cv-form/models"
)

type EmploymentRepository interface {
	CreateEmployment(employment *models.Employment) (models.Employment, error)
	GetEmployments(ProfileCode uint) ([]models.Employment, error)
	FindEmployments(id uint) (models.Employment, error)
	DeleteEmployment(id uint) error
}

type employmentRepository struct{}

func NewEmploymentRepository() EmploymentRepository {
	return &employmentRepository{}
}

func (repo *employmentRepository) CreateEmployment(employment *models.Employment) (models.Employment, error) {
	err := config.DB.Create(&employment).Error
	if err != nil {
		return *employment, err
	}
	return *employment, nil
}

func (repo *employmentRepository) GetEmployments(ProfileCode uint) ([]models.Employment, error) {
	var employments []models.Employment
	err := config.DB.Where("profile_code = ?", ProfileCode).Find(&employments).Error
	if err != nil {
		return nil, err
	}
	return employments, nil
}

func (repo *employmentRepository) FindEmployments(id uint) (models.Employment, error) {
	var employment models.Employment
	err := config.DB.Where("id =?", id).First(&employment).Error
	if err != nil {
		return employment, err
	}
	return employment, nil
}

func (repo *employmentRepository) DeleteEmployment(id uint) error {
	return config.DB.Delete(&models.Employment{ID: id}).Error
}
