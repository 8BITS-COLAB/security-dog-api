package dtos

import "errors"

type SigninDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RemoteIP string `json:"remote_ip"`
	Geo      struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	} `json:"geo"`
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
