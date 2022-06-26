package entities

type Device struct {
	Base
	UserID    string  `json:"user_id" gorm:"uniqueIndex:idx_remote_addr" `
	User      User    `json:"user" gorm:"foreignKey:user_id"`
	Name      string  `json:"name"`
	RemoteIP  string  `json:"remote_ip" gorm:"uniqueIndex:idx_remote_addr"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	IsLinked  bool    `json:"is_linked"`
	IsTrusted bool    `json:"is_trusted"`
	IsBlocked bool    `json:"is_blocked"`
}
