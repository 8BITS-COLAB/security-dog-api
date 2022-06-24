package services

import (
	"errors"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"github.com/ElioenaiFerrari/security-dog-api/views"
	"github.com/andskur/argon2-hashing"
	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
)

type SharedRegistryService struct {
	db              *cache.Cache
	registryService *RegistryService
}

func NewSharedRegistryService(db *cache.Cache, registryService *RegistryService) *SharedRegistryService {
	return &SharedRegistryService{db: db, registryService: registryService}
}

func (sharedRegistryService *SharedRegistryService) Create(createSharedRegistryDTO *dtos.CreateSharedRegistryDTO) (entities.SharedRegistry, error) {
	var sharedRegistry entities.SharedRegistry

	hash, err := argon2.GenerateFromPassword([]byte(createSharedRegistryDTO.Password), argon2.DefaultParams)

	if err != nil {
		return sharedRegistry, err
	}

	sharedRegistry.ID = uuid.NewString()
	sharedRegistry.UserID = createSharedRegistryDTO.UserID
	sharedRegistry.RegistryID = createSharedRegistryDTO.RegistryID
	sharedRegistry.ExpireAt = createSharedRegistryDTO.ExpireAt
	sharedRegistry.Password = string(hash)

	if err := sharedRegistryService.db.Add(sharedRegistry.ID, sharedRegistry, sharedRegistry.ExpireAt); err != nil {
		return sharedRegistry, err
	}

	return sharedRegistry, nil
}

func (sharedRegistryService *SharedRegistryService) GetByID(id, password string) (views.RegistryView, error) {
	var registryView views.RegistryView
	sharedRegistry, _ := sharedRegistryService.db.Get(id)

	if sharedRegistry == nil {
		return registryView, errors.New("shared registry not found")
	}

	if err := argon2.CompareHashAndPassword([]byte(sharedRegistry.(entities.SharedRegistry).Password), []byte(password)); err != nil {
		return registryView, errors.New("invalid password")
	}

	registryView, err := sharedRegistryService.registryService.GetByID(sharedRegistry.(entities.SharedRegistry).UserID, sharedRegistry.(entities.SharedRegistry).RegistryID)

	if err != nil {
		return registryView, err
	}

	return registryView, nil
}
