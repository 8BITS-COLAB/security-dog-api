package security

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	UserName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
