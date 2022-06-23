package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/security"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"github.com/golang-jwt/jwt"
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

	token, err := authController.authService.Signin(&signinDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (authController *AuthController) Signout(c echo.Context) error {
	currentUser := c.Get("user").(*jwt.Token)
	claims := currentUser.Claims.(*security.JwtClaims)

	remoteIP := c.QueryParam("remote_ip")

	if remoteIP == "" {
		remoteIP = c.RealIP()
	}

	if err := authController.authService.Signout(claims.Subject, remoteIP); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (authController *AuthController) Profile(c echo.Context) error {
	currentUser := c.Get("user").(*jwt.Token)
	claims := currentUser.Claims.(*security.JwtClaims)

	user, err := authController.authService.Profile(claims.Subject)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}
