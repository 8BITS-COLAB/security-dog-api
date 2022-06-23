package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/security"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"github.com/labstack/echo/v4"
)

type RegistryController struct {
	registryService *services.RegistryService
}

func NewRegistryController(registryService *services.RegistryService) *RegistryController {
	return &RegistryController{registryService: registryService}
}

func (registryController *RegistryController) Create(c echo.Context) error {
	var createRegistryDTO dtos.CreateRegistryDTO

	if err := c.Bind(&createRegistryDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createRegistryDTO.UserID = c.Param("user_id")

	if err := createRegistryDTO.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	registry, err := registryController.registryService.Create(&createRegistryDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, registry)
}

func (registryController *RegistryController) Index(c echo.Context) error {
	claims := security.Claims(c)

	userID := c.Param("user_id")

	if claims.Subject != userID {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to see this registries")
	}

	registries, err := registryController.registryService.GetAll(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, registries)
}

func (registryController *RegistryController) Show(c echo.Context) error {
	claims := security.Claims(c)

	userID := c.Param("user_id")
	id := c.Param("id")

	if claims.Subject != userID {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to see this registry")
	}

	registry, err := registryController.registryService.GetByID(userID, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, registry)
}

func (registryController *RegistryController) Update(c echo.Context) error {
	claims := security.Claims(c)

	userID := c.Param("user_id")
	id := c.Param("id")

	var updateRegistryDTO dtos.UpdateRegistryDTO

	if err := c.Bind(&updateRegistryDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := updateRegistryDTO.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if claims.Subject != userID {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to update this registry")
	}

	registry, err := registryController.registryService.Update(userID, id, &updateRegistryDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, registry)
}

func (registryController *RegistryController) Delete(c echo.Context) error {
	claims := security.Claims(c)

	userID := c.Param("user_id")
	id := c.Param("id")

	if claims.Subject != userID {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to delete this registry")
	}

	err := registryController.registryService.Delete(userID, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
