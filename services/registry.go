package services

import (
	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"gorm.io/gorm"
)

type RegistryService struct {
	db *gorm.DB
}

func NewRegistryService(db *gorm.DB) *RegistryService {
	return &RegistryService{db: db}
}

func (registryService *RegistryService) Create(createRegistryDTO *dtos.CreateRegistryDTO) (entities.Registry, error) {
	registry := entities.Registry{
		UserID:   createRegistryDTO.UserID,
		Name:     createRegistryDTO.Name,
		Login:    createRegistryDTO.Login,
		Password: createRegistryDTO.Password,
	}

	if err := registryService.db.Create(&registry).Error; err != nil {
		return registry, err
	}

	return registry, nil
}

func (registryService *RegistryService) GetAll(userID string) ([]entities.Registry, error) {
	var registries []entities.Registry

	if err := registryService.db.Where("user_id = ?", userID).Find(&registries).Error; err != nil {
		return registries, err
	}

	return registries, nil
}

func (registryService *RegistryService) GetByID(userID, id string) (entities.Registry, error) {
	var registry entities.Registry

	if err := registryService.db.Where("user_id = ? AND id = ?", userID, id).First(&registry).Error; err != nil {
		return registry, err
	}

	return registry, nil
}

func (registryService *RegistryService) Update(userID, id string, updateRegistryDTO *dtos.UpdateRegistryDTO) (entities.Registry, error) {
	var registry entities.Registry

	if err := registryService.db.Where("user_id = ? AND id = ?", userID, id).First(&registry).Error; err != nil {
		return registry, err
	}

	registry.Name = updateRegistryDTO.Name
	registry.Login = updateRegistryDTO.Login
	registry.Password = updateRegistryDTO.Password

	if err := registryService.db.Updates(&registry).Error; err != nil {
		return registry, err
	}

	return registry, nil
}

func (registryService *RegistryService) Delete(userID, id string) error {
	if err := registryService.db.Where("user_id = ? AND id = ?", userID, id).Delete(&entities.Registry{}).Error; err != nil {
		return err
	}

	return nil
}
