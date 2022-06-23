package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/services"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (userController *UserController) Index(c echo.Context) error {
	users, err := userController.userService.GetAll()

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
