package services

import (
	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"github.com/ElioenaiFerrari/security-dog-api/views"
	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (userService *UserService) Create(createUserDTO *dtos.CreateUserDTO) (views.UserView, error) {
	var userView views.UserView

	user := entities.User{
		UserName: createUserDTO.UserName,
		Password: createUserDTO.Password,
		Email:    createUserDTO.Email,
		Role:     createUserDTO.Role,
	}

	if err := userService.db.Create(&user).Scan(&userView).Error; err != nil {
		return userView, err
	}

	return userView, nil
}

func (userService *UserService) GetAll() ([]views.UserView, error) {
	var users []views.UserView

	if err := userService.db.Model(entities.User{}).Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func (userService *UserService) GetByID(id string) (views.UserView, string, error) {
	var user entities.User
	var userView views.UserView

	if err := userService.db.First(&user, "id = ?", id).Scan(&userView).Error; err != nil {
		return userView, user.SecretKey, err
	}

	return userView, user.SecretKey, nil
}

func (userService *UserService) GetByEmail(email string) (entities.User, error) {
	var user entities.User

	if err := userService.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (userService *UserService) Update(id string, updateUserDTO *dtos.UpdateUserDTO) (views.UserView, error) {
	var userView views.UserView
	var user entities.User

	if err := userService.db.Where("id = ?", id).First(&user).Error; err != nil {
		return userView, err
	}

	user.UserName = updateUserDTO.UserName
	user.Password = updateUserDTO.Password
	user.Email = updateUserDTO.Email
	user.Role = updateUserDTO.Role

	if err := userService.db.UpdateColumns(&user).Scan(&userView).Error; err != nil {
		return userView, err
	}

	return userView, nil
}

func (userService *UserService) Delete(id string) error {
	var user entities.User

	if _, _, err := userService.GetByID(id); err != nil {
		return err
	}

	if err := userService.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
