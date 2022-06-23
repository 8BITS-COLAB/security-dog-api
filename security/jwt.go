package security

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenToken(subject string) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		Subject:   subject,
		Issuer:    os.Getenv("JWT_ISSUER"),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
