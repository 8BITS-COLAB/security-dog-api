package views

type UserView struct {
	BaseView
	UserName string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
