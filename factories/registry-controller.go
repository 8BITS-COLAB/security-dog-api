package factories

import (
	"github.com/ElioenaiFerrari/security-dog-api/controllers"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"gorm.io/gorm"
)

func MakeRegistryController(db *gorm.DB) *controllers.RegistryController {
	registryService := services.NewRegistryService(db)

	return controllers.NewRegistryController(registryService)
}
