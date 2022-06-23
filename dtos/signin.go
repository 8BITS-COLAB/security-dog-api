package dtos

import "errors"

type SigninDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (signinDTO SigninDTO) Validate() error {
	if signinDTO.Email == "" {
		return errors.New("email is required")
	}

	if signinDTO.Password == "" {
		return errors.New("password is required")
	}

	return nil
}