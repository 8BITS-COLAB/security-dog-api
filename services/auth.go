package services

import (
	"errors"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"github.com/ElioenaiFerrari/security-dog-api/security"
	"github.com/ElioenaiFerrari/security-dog-api/views"
	"github.com/andskur/argon2-hashing"
	"gorm.io/gorm"
)

type AuthService struct {
	db            *gorm.DB
	userService   *UserService
	deviceService *DeviceService
}

func NewAuthService(db *gorm.DB, userService *UserService, deviceService *DeviceService) *AuthService {
	return &AuthService{
		db:            db,
		userService:   userService,
		deviceService: deviceService,
	}
}

func (authService *AuthService) Signup(createUserDTO *dtos.CreateUserDTO) (views.UserView, error) {
	user, err := authService.userService.Create(createUserDTO)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (authService *AuthService) Signin(signinDTO *dtos.SigninDTO) (string, entities.User, error) {
	user, err := authService.userService.GetByEmail(signinDTO.Email)

	if err != nil {
		return "", user, err
	}

	if err := argon2.CompareHashAndPassword([]byte(user.Password), []byte(signinDTO.Password)); err != nil {
		return "", user, errors.New("invalid password")
	}

	if _, err = authService.deviceService.Add(user.ID, signinDTO.RemoteIP); err != nil {
		return "", user, err
	}

	accessToken, err := security.GenToken(user.UserName, user.Email, user.Role, user.ID)

	if err != nil {
		return "", user, err
	}

	return accessToken, user, nil
}

func (authService *AuthService) Profile(userID string) (views.UserView, error) {
	user, err := authService.userService.GetByID(userID)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (authService *AuthService) ValidateDevice(userID string, remoteIP string) error {
	device, err := authService.deviceService.GetByRemoteIP(userID, remoteIP)

	if err != nil {
		return err
	}

	if !device.IsLinked {
		return errors.New("device is not linked")
	}

	if !device.IsTrusted {
		return errors.New("device is not trusted")
	}

	if device.IsBlocked {
		return errors.New("device is blocked")
	}

	return nil
}
