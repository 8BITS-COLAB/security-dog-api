package services

import (
	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"github.com/andskur/argon2-hashing"
	"gorm.io/gorm"
)

type AuthService struct {
	db          *gorm.DB
	userService *UserService
}

func NewAuthService(db *gorm.DB, userService *UserService) *AuthService {
	return &AuthService{db: db, userService: userService}
}

func (authService *AuthService) Signup(createUserDTO *dtos.CreateUserDTO) (entities.User, error) {
	user, err := authService.userService.Create(createUserDTO)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (authService *AuthService) Signin(signinDTO *dtos.SigninDTO) (entities.User, error) {
	user, err := authService.userService.GetByEmail(signinDTO.Email)

	if err != nil {
		return user, err
	}

	if err := argon2.CompareHashAndPassword([]byte(user.Password), []byte(signinDTO.Password)); err != nil {
		return user, err
	}

	return user, nil
}
