package entities

import (
	"github.com/ElioenaiFerrari/security-dog-api/security"
	"gorm.io/gorm"
)

type Registry struct {
	Base
	UserID   string `json:"user_id" gorm:"index"`
	User     User   `json:"user" gorm:"foreignKey:user_id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	SiteURL  string `json:"site_url"`
}

func (registry *Registry) BeforeCreate(tx *gorm.DB) (err error) {
	var user User

	tx.Find(&user, "id = ?", registry.UserID)

	privateKey, err := security.PrivateKeyFromString(user.PrivateKey)

	if err != nil {
		return err
	}

	cipherText, err := security.Encrypt(&privateKey.PublicKey, registry.Password)

	if err != nil {
		return err
	}

	registry.Password = cipherText

	return nil
}

func (registry *Registry) AfterFind(tx *gorm.DB) error {
	var user User

	tx.Find(&user, "id = ?", registry.UserID)

	privateKey, err := security.PrivateKeyFromString(user.PrivateKey)

	if err != nil {
		return err
	}

	text, err := security.Decrypt(privateKey, registry.Password)

	if err != nil {
		return err
	}

	registry.Password = text

	return nil
}
