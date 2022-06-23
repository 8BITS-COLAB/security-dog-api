package factories

import (
	"github.com/ElioenaiFerrari/security-dog-api/controllers"
	"github.com/ElioenaiFerrari/security-dog-api/services"
	"gorm.io/gorm"
)

func MakeUserController(db *gorm.DB) *controllers.UserController {
	userService := services.NewUserService(db)

	return controllers.NewUserController(userService)
}
