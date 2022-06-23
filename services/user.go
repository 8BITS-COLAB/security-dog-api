package services

import (
	"fmt"

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

func (userService *UserService) GetByID(id string) (entities.User, error) {
	var user entities.User

	if err := userService.db.First(&user, "id = ?", id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (userService *UserService) GetByEmail(email string) (entities.User, error) {
	var user entities.User

	if err := userService.db.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (userService *UserService) Update(id string, updateUserDTO *dtos.UpdateUserDTO) (entities.User, error) {
	var user entities.User

	if _, err := userService.GetByID(id); err != nil {
		return user, err
	}

	user.UserName = updateUserDTO.UserName
	user.Password = updateUserDTO.Password
	user.Email = updateUserDTO.Email
	user.Role = updateUserDTO.Role

	if err := userService.db.Model(user).Where("id = ?", id).Updates(&user).Error; err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}

func (userService *UserService) Delete(id string) error {
	var user entities.User

	if _, err := userService.GetByID(id); err != nil {
		return err
	}

	if err := userService.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
