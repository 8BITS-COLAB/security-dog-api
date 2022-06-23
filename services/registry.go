package services

import (
	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"github.com/ElioenaiFerrari/security-dog-api/views"
	"gorm.io/gorm"
)

type RegistryService struct {
	db *gorm.DB
}

func NewRegistryService(db *gorm.DB) *RegistryService {
	return &RegistryService{db: db}
}

func (registryService *RegistryService) Create(createRegistryDTO *dtos.CreateRegistryDTO) (views.RegistryView, error) {
	var registryView views.RegistryView

	if err := registryService.db.Create(&entities.Registry{
		UserID:   createRegistryDTO.UserID,
		Name:     createRegistryDTO.Name,
		Login:    createRegistryDTO.Login,
		Password: createRegistryDTO.Password,
		SiteURL:  createRegistryDTO.SiteURL,
	}).Scan(&registryView).Error; err != nil {
		return registryView, err
	}

	return registryView, nil
}

func (registryService *RegistryService) GetAll(userID string) ([]views.RegistryView, error) {
	var registriesView []views.RegistryView

	if err := registryService.db.Where("user_id = ?", userID).Find(&entities.Registry{}).Scan(&registriesView).Error; err != nil {
		return registriesView, err
	}

	return registriesView, nil
}

func (registryService *RegistryService) GetByID(userID, id string) (views.RegistryView, error) {
	var registryView views.RegistryView

	if err := registryService.db.Where("user_id = ? AND id = ?", userID, id).First(&entities.Registry{}).Scan(&registryView).Error; err != nil {
		return registryView, err
	}

	return registryView, nil
}

func (registryService *RegistryService) Update(userID, id string, updateRegistryDTO *dtos.UpdateRegistryDTO) (views.RegistryView, error) {
	var registry entities.Registry
	var registryView views.RegistryView

	if err := registryService.db.Where("user_id = ? AND id = ?", userID, id).First(&registry).Error; err != nil {
		return registryView, err
	}

	registry.Name = updateRegistryDTO.Name
	registry.Login = updateRegistryDTO.Login
	registry.Password = updateRegistryDTO.Password
	registry.SiteURL = updateRegistryDTO.SiteURL

	if err := registryService.db.Updates(&registry).Scan(&registryView).Error; err != nil {
		return registryView, err
	}

	return registryView, nil
}

func (registryService *RegistryService) Delete(userID, id string) error {
	if err := registryService.db.Where("user_id = ? AND id = ?", userID, id).Delete(&entities.Registry{}).Error; err != nil {
		return err
	}

	return nil
}
