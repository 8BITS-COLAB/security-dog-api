package dtos

import (
	"errors"
	"time"
)

type CreateSharedRegistryDTO struct {
	UserID     string        `json:"user_id"`
	RegistryID string        `json:"registry_id"`
	ExpireAt   time.Duration `json:"expire_at"`
	Password   string        `json:"password"`
}

func (createSharedRegistryDTO *CreateSharedRegistryDTO) Validate() error {
	if createSharedRegistryDTO.UserID == "" {
		return errors.New("user_id is required")
	}

	if createSharedRegistryDTO.RegistryID == "" {
		return errors.New("registry_id is required")
	}

	if createSharedRegistryDTO.ExpireAt == 0 {
		return errors.New("expire_at is required")
	}

	if createSharedRegistryDTO.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
