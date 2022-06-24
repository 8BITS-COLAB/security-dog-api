package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/security"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	accessToken, user, err := authController.authService.Signin(&signinDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := authController.authService.ValidateDevice(user.ID, signinDTO.RemoteIP); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token": accessToken,
	})
}

func (authController *AuthController) Profile(c echo.Context) error {
	claims := security.Claims(c)

	user, err := authController.authService.Profile(claims.Subject)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (authController *AuthController) RefreshToken(c echo.Context) error {
	claims := security.Claims(c)

	remoteIP := c.RealIP()

	if err := authController.authService.ValidateDevice(claims.Subject, remoteIP); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	accessToken, err := security.GenToken(claims.UserName, claims.Email, claims.Role, claims.Subject)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"access_token": accessToken,
	})
}

func (authController *AuthController) CSRFToken(c echo.Context) error {
	csrf := c.Get(middleware.DefaultCSRFConfig.ContextKey)

	return c.JSON(http.StatusOK, map[string]string{
		"csrf_token": csrf.(string),
	})

}
