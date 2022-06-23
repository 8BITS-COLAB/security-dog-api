package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (authController *AuthController) Signup(c echo.Context) error {
	var createUserDTO dtos.CreateUserDTO

	if err := c.Bind(&createUserDTO); err != nil {
		return err
	}

	if err := createUserDTO.Validate(); err != nil {
		return err
	}

	user, err := authController.authService.Signup(&createUserDTO)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func (authController *AuthController) Signin(c echo.Context) error {
	var signinDTO dtos.SigninDTO

	if err := c.Bind(&signinDTO); err != nil {
		return err
	}

	if err := signinDTO.Validate(); err != nil {
		return err
	}

	user, err := authController.authService.Signin(&signinDTO)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)

}
