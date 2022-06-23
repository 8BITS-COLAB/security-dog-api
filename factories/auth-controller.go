package factories

import (
	"github.com/ElioenaiFerrari/security-dog-api/controllers"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"gorm.io/gorm"
)

func MakeAuthController(db *gorm.DB) *controllers.AuthController {
	userService := services.NewUserService(db)
	return controllers.NewAuthController(userService)
}
