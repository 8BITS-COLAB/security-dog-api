package dtos

import (
	"errors"
)

type CreateSharedRegistryDTO struct {
	UserID     string `json:"user_id"`
	RegistryID string `json:"registry_id"`
	ExpireAt   string `json:"expire_at"`
	Password   string `json:"password"`
}

func (createSharedRegistryDTO *CreateSharedRegistryDTO) Validate() error {
	if createSharedRegistryDTO.UserID == "" {
		return errors.New("user_id is required")
	}

	if createSharedRegistryDTO.RegistryID == "" {
		return errors.New("registry_id is required")
	}

	if createSharedRegistryDTO.ExpireAt == "" || createSharedRegistryDTO.ExpireAt != "1m" && createSharedRegistryDTO.ExpireAt != "5m" && createSharedRegistryDTO.ExpireAt != "30m" && createSharedRegistryDTO.ExpireAt != "1h" && createSharedRegistryDTO.ExpireAt != "6h" && createSharedRegistryDTO.ExpireAt != "1d" {
		return errors.New("invalid expire_at. Only 1m, 5m, 30m, 1h, 6h, 1d are allowed")
	}

	if createSharedRegistryDTO.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
