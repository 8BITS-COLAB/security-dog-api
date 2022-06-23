package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	userService *services.UserService
}

func NewAuthController(userService *services.UserService) *AuthController {
	return &AuthController{userService: userService}
}

func (authController *AuthController) Signup(c echo.Context) error {
	var createUserDTO dtos.CreateUserDTO

	if err := c.Bind(&createUserDTO); err != nil {
		return err
	}

	if err := createUserDTO.Validate(); err != nil {
		return err
	}

	user, err := authController.userService.Create(&createUserDTO)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}
