package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"github.com/labstack/echo/v4"
)

type SharedRegistryController struct {
	sharedRegistryService *services.SharedRegistryService
}

func NewSharedRegistryController(sharedRegistryService *services.SharedRegistryService) *SharedRegistryController {
	return &SharedRegistryController{sharedRegistryService: sharedRegistryService}
}

func (sharedRegistryController *SharedRegistryController) Create(c echo.Context) error {
	var createSharedRegistryDTO dtos.CreateSharedRegistryDTO

	if err := c.Bind(&createSharedRegistryDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createSharedRegistryDTO.UserID = c.Param("user_id")

	if err := createSharedRegistryDTO.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sharedRegistry, err := sharedRegistryController.sharedRegistryService.Create(&createSharedRegistryDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, sharedRegistry)
}

func (sharedRegistryController *SharedRegistryController) Show(c echo.Context) error {
	ID := c.Param("id")
	password := c.QueryParam("password")

	registry, err := sharedRegistryController.sharedRegistryService.GetByID(ID, password)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, registry)
}
