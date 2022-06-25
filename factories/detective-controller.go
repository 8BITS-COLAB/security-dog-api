package factories

import (
	"github.com/ElioenaiFerrari/security-dog-api/controllers"
	"github.com/ElioenaiFerrari/security-dog-api/services"
)

func MakeDetectiveController() *controllers.DetectiveController {
	detectiveService := services.NewDetectiveService()

	return controllers.NewDetectiveController(detectiveService)
}
