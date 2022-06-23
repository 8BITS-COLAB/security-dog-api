package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"github.com/labstack/echo/v4"
)

type DeviceController struct {
	deviceService *services.DeviceService
}

func NewDeviceController(deviceService *services.DeviceService) *DeviceController {
	return &DeviceController{deviceService: deviceService}
}

func (deviceController *DeviceController) Index(c echo.Context) error {
	userID := c.Param("user_id")

	devices, err := deviceController.deviceService.GetAll(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, devices)
}

func (deviceController *DeviceController) Update(c echo.Context) error {
	var updateDeviceDTO dtos.UpdateDeviceDTO

	if err := c.Bind(&updateDeviceDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updateDeviceDTO.UserID = c.Param("user_id")

	device, err := deviceController.deviceService.Update(&updateDeviceDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, device)
}
