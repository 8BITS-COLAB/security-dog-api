package dtos

type UpdateRegistryDTO struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	SiteURL  string `json:"site_url"`
}

func (updateRegistryDTO *UpdateRegistryDTO) Validate() error {

	return nil
}
