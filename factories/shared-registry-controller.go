package factories

import (
	"github.com/ElioenaiFerrari/security-dog-api/controllers"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"gorm.io/gorm"
)

func MakeSharedRegistryController(db *gorm.DB) *controllers.SharedRegistryController {
	registryService := services.NewRegistryService(db)
	sharedRegistryService := services.NewSharedRegistryService(registryService)

	return controllers.NewSharedRegistryController(sharedRegistryService)

}
