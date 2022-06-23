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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := createUserDTO.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := authController.authService.Signup(&createUserDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (authController *AuthController) Signin(c echo.Context) error {
	var signinDTO dtos.SigninDTO

	if err := c.Bind(&signinDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	signinDTO.RemoteIP = c.RealIP()

	if err := signinDTO.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := authController.authService.Signin(&signinDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)

}
