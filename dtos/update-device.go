package dtos

type UpdateDeviceDTO struct {
	UserID    string `json:"user_id"`
	RemoteIP  string `json:"remote_ip"`
	IsLinked  bool   `json:"is_linked"`
	IsTrusted bool   `json:"is_trusted"`
	IsBlocked bool   `json:"is_blocked"`
}

func (updateDeviceDTO *UpdateDeviceDTO) Validate() error {
	return nil
}
