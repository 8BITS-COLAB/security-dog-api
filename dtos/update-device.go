package dtos

type UpdateDeviceDTO struct {
	UserID    string `json:"user_id"`
	ID        string `json:"id"`
	IsLinked  bool   `json:"is_linked"`
	IsTrusted bool   `json:"is_trusted"`
	IsBlocked bool   `json:"is_blocked"`
}
