package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
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

func (userController *UserController) Show(c echo.Context) error {
	id := c.Param("id")

	user, err := userController.userService.GetByID(id)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (userController *UserController) Update(c echo.Context) error {
	id := c.Param("id")

	var updateUserDTO dtos.UpdateUserDTO

	if err := c.Bind(&updateUserDTO); err != nil {
		return err
	}

	if err := updateUserDTO.Validate(); err != nil {
		return err
	}

	user, err := userController.userService.Update(id, &updateUserDTO)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (userController *UserController) Delete(c echo.Context) error {
	id := c.Param("id")

	err := userController.userService.Delete(id)

	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
