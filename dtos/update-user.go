package dtos

import "errors"

type UpdateUserDTO struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (updateUserDTO UpdateUserDTO) Validate() error {
	if updateUserDTO.Role != "" && updateUserDTO.Role != "admin" && updateUserDTO.Role != "user" {
		return errors.New("role must be admin or user")
	}

	return nil
}
