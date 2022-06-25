package controllers

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/services"
	"github.com/labstack/echo/v4"
)

type DetectiveController struct {
	detectiveService *services.DetectiveService
}

func NewDetectiveController(detectiveService *services.DetectiveService) *DetectiveController {
	return &DetectiveController{detectiveService: detectiveService}
}

func (detectiveController *DetectiveController) Investigate(c echo.Context) error {
	key := c.QueryParam("key")
	result, err := detectiveController.detectiveService.Investigate(key)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, result)
}
