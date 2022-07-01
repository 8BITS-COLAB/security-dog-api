package entities

import (
	"crypto/rand"
	"crypto/rsa"

	"github.com/ElioenaiFerrari/security-dog-api/security"
	"github.com/andskur/argon2-hashing"
	"github.com/xlzd/gotp"
	"gorm.io/gorm"
)

type User struct {
	Base

	UserName   string     `json:"username"`
	Email      string     `json:"email" gorm:"uniqueIndex"`
	Password   string     `json:"password"`
	Role       string     `json:"role"`
	PrivateKey string     `json:"private_key"`
	SecretKey  string     `json:"secret_key"`
	Registries []Registry `json:"registries" gorm:"foreignKey:user_id"`
	Devices    []Device   `json:"devices" gorm:"foreignKey:user_id"`
	Has2FA     bool       `json:"has_2fa"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := argon2.GenerateFromPassword([]byte(user.Password), argon2.DefaultParams)

	if err != nil {
		return err
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		return err
	}

	block, err := security.PrivateKeyToString(privateKey)

	if err != nil {
		return err
	}

	user.Password = string(hash)
	user.PrivateKey = string(block)
	user.SecretKey = gotp.RandomSecret(16)

	return err
}
