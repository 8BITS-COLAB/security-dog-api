package views

type RegistryView struct {
	BaseView
	UserID   string `json:"user_id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	SiteURL  string `json:"site_url"`
}
