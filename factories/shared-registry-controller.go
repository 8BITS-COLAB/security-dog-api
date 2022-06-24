package factories

import (
	"github.com/ElioenaiFerrari/security-dog-api/controllers"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

func MakeSharedRegistryController(memory *cache.Cache, db *gorm.DB) *controllers.SharedRegistryController {
	registryService := services.NewRegistryService(db)
	sharedRegistryService := services.NewSharedRegistryService(memory, registryService)

	return controllers.NewSharedRegistryController(sharedRegistryService)

}
