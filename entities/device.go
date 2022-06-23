package entities

type Device struct {
	Base
	UserID    string `json:"user_id" gorm:"index"`
	User      User   `json:"user" gorm:"foreignKey:user_id"`
	Name      string `json:"name"`
	RemoteIP  string `json:"remote_ip"`
	IsLinked  bool   `json:"is_linked"`
	IsTrusted bool   `json:"is_trusted"`
	IsBlocked bool   `json:"is_blocked"`
}
