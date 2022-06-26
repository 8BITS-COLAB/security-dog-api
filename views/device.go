package views

type DeviceView struct {
	BaseView
	UserID    string  `json:"user_id"`
	RemoteIP  string  `json:"remote_ip"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
	IsLinked  bool    `json:"is_linked"`
	IsTrusted bool    `json:"is_trusted"`
	IsBlocked bool    `json:"is_blocked"`
}
