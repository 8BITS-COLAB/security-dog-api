package entities

import "time"

type SharedRegistry struct {
	ID         string        `json:"id"`
	UserID     string        `json:"user_id"`
	RegistryID string        `json:"registry_id"`
	ExpireAt   time.Duration `json:"expire_at"`
	Password   string        `json:"password"`
}
