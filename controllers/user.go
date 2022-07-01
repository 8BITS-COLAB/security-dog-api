package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/security"
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
	claims := security.Claims(c)

	if claims.Role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to list users")
	}

	users, err := userController.userService.GetAll()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}

func (userController *UserController) Show(c echo.Context) error {
	claims := security.Claims(c)

	id := c.Param("id")

	if claims.Role != "admin" && claims.Subject != id {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to see this user")
	}

	user, _, err := userController.userService.GetByID(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (userController *UserController) Update(c echo.Context) error {
	claims := security.Claims(c)

	id := c.Param("id")

	if claims.Subject != id {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to update this user")
	}

	var updateUserDTO dtos.UpdateUserDTO

	if err := c.Bind(&updateUserDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := updateUserDTO.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := userController.userService.Update(id, &updateUserDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (userController *UserController) Delete(c echo.Context) error {
	claims := security.Claims(c)

	id := c.Param("id")

	if claims.Subject != id {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to delete this user")
	}

	err := userController.userService.Delete(id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
