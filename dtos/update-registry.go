package dtos

type UpdateRegistryDTO struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (updateRegistryDTO *UpdateRegistryDTO) Validate() error {

	return nil
}
