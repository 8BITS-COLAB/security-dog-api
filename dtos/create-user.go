package dtos

import "errors"

type CreateUserDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (createUserDTO CreateUserDTO) Validate() error {

	if createUserDTO.UserName == "" {
		return errors.New("username is required")
	}

	if createUserDTO.Password == "" {
		return errors.New("password is required")
	}

	if createUserDTO.Email == "" {
		return errors.New("email is required")
	}

	if createUserDTO.Role == "" {
		return errors.New("role is required")
	}

	if createUserDTO.Role != "admin" && createUserDTO.Role != "user" {
		return errors.New("role must be admin or user")
	}

	return nil
}
