package dtos

import (
	"errors"
)

type CreateRegistryDTO struct {
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (createRegistryDTO *CreateRegistryDTO) Validate() error {
	if createRegistryDTO.UserID == "" {
		return errors.New("user_id is required")
	}

	if createRegistryDTO.Name == "" {
		return errors.New("name is required")
	}

	if createRegistryDTO.Login == "" {
		return errors.New("login is required")
	}

	if createRegistryDTO.Password == "" {
		return errors.New("password is required")
	}

	return nil
}
