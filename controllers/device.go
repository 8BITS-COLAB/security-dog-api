package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/dtos"
	"github.com/ElioenaiFerrari/security-dog-api/security"
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
	claims := security.Claims(c)

	userID := c.Param("user_id")

	if claims.Subject != userID {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to see this devices")
	}

	devices, err := deviceController.deviceService.GetAll(userID)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, devices)
}

func (deviceController *DeviceController) Update(c echo.Context) error {
	claims := security.Claims(c)

	var updateDeviceDTO dtos.UpdateDeviceDTO

	if err := c.Bind(&updateDeviceDTO); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	updateDeviceDTO.UserID = c.Param("user_id")
	updateDeviceDTO.RemoteIP = c.QueryParam("remote_ip")

	if err := updateDeviceDTO.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if claims.Subject != updateDeviceDTO.UserID {
		return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to update this device")
	}

	device, err := deviceController.deviceService.Update(&updateDeviceDTO)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, device)
}
