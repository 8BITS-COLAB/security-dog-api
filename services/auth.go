package services

import (
	"errors"
	"os"
	"time"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"github.com/ElioenaiFerrari/security-dog-api/security"
	"github.com/andskur/argon2-hashing"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type AuthService struct {
	db            *gorm.DB
	userService   *UserService
	deviceService *DeviceService
}

func NewAuthService(db *gorm.DB, userService *UserService, deviceService *DeviceService) *AuthService {
	return &AuthService{db: db, userService: userService, deviceService: deviceService}
}

func (authService *AuthService) Signup(createUserDTO *dtos.CreateUserDTO) (entities.User, error) {
	user, err := authService.userService.Create(createUserDTO)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (authService *AuthService) Signin(signinDTO *dtos.SigninDTO) (string, error) {
	user, err := authService.userService.GetByEmail(signinDTO.Email)

	if err != nil {
		return "", err
	}

	if err := argon2.CompareHashAndPassword([]byte(user.Password), []byte(signinDTO.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	if _, err = authService.deviceService.Add(user.ID, signinDTO.RemoteIP); err != nil {
		return "", err
	}

	claims := &security.JwtClaims{
		UserName: user.UserName,
		Role:     user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Subject:   user.ID,
			Issuer:    os.Getenv("JWT_ISSUER"),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (authService *AuthService) Signout(userID string, remoteIP string) error {
	_, err := authService.deviceService.Update(&dtos.UpdateDeviceDTO{
		UserID:   userID,
		RemoteIP: remoteIP,
		IsLinked: false,
	})

	if err != nil {
		return err
	}

	return nil
}

func (authService *AuthService) Profile(userID string) (entities.User, error) {
	user, err := authService.userService.GetByID(userID)

	if err != nil {
		return user, err
	}

	return user, nil
}
