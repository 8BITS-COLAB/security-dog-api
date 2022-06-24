package dtos

import "errors"

type UpdateDeviceDTO struct {
	UserID    string `json:"user_id"`
	RemoteIP  string `json:"remote_ip"`
	IsLinked  bool   `json:"is_linked"`
	IsTrusted bool   `json:"is_trusted"`
	IsBlocked bool   `json:"is_blocked"`
}

func (updateDeviceDTO *UpdateDeviceDTO) Validate() error {

	if updateDeviceDTO.UserID == "" {
		return errors.New("user_id is required")
	}

	if updateDeviceDTO.RemoteIP == "" {
		return errors.New("remote_ip is required")
	}

	return nil
}
