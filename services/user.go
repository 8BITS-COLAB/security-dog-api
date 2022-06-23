package services

import (
	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (userService *UserService) Create(createUserDTO *dtos.CreateUserDTO) (entities.User, error) {
	user := entities.User{
		UserName: createUserDTO.UserName,
		Password: createUserDTO.Password,
		Email:    createUserDTO.Email,
		Role:     createUserDTO.Role,
	}

	if err := userService.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (userService *UserService) GetAll() ([]entities.User, error) {
	var users []entities.User

	if err := userService.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
