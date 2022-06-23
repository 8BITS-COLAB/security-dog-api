package factories

import (
	"github.com/ElioenaiFerrari/security-dog-api/controllers"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"gorm.io/gorm"
)

func MakeDeviceController(db *gorm.DB) *controllers.DeviceController {
	deviceService := services.NewDeviceService(db)

	return controllers.NewDeviceController(deviceService)
}
