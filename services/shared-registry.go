package services

import (
	"errors"
	"fmt"
	"time"

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
	sharedRegistry.Password = string(hash)

	switch createSharedRegistryDTO.ExpireAt {
	case "1m":
		sharedRegistry.ExpireAt = time.Minute
	case "5m":
		sharedRegistry.ExpireAt = time.Minute * 5
	case "30m":
		sharedRegistry.ExpireAt = time.Minute * 30
	case "1h":
		sharedRegistry.ExpireAt = time.Hour
	case "6h":
		sharedRegistry.ExpireAt = time.Hour * 6
	case "1d":
		sharedRegistry.ExpireAt = time.Hour * 24
	default:
		sharedRegistry.ExpireAt = time.Minute * 5
	}

	fmt.Println(sharedRegistry)

	sharedRegistryService.db.Set(sharedRegistry.ID, sharedRegistry, sharedRegistry.ExpireAt)

	return sharedRegistry, nil
}

func (sharedRegistryService *SharedRegistryService) GetByID(id, password string) (views.RegistryView, error) {
	var registryView views.RegistryView
	sharedRegistry, found := sharedRegistryService.db.Get(id)

	if !found {
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
