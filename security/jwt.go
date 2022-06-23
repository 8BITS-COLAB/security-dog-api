package security

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JwtClaims struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenToken(userName, email, role, subject string) (string, error) {
	claims := &JwtClaims{
		UserName: userName,
		Email:    email,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Subject:   subject,
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Claims(c echo.Context) *JwtClaims {
	accessToken := c.Get("user").(*jwt.Token)
	return accessToken.Claims.(*JwtClaims)

}
