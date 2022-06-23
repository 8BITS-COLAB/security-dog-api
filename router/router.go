package router

import (
	"net/http"

	"github.com/ElioenaiFerrari/security-dog-api/factories"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Route struct {
	Func        echo.HandlerFunc
	Path        string
	Middlewares []echo.MiddlewareFunc
	Method      string
}

func InitV1(v1 *echo.Group, db *gorm.DB) {
	authController := factories.MakeAuthController(db)
	userController := factories.MakeUserController(db)

	var Routes = []Route{
		{
			Func:        authController.Signup,
			Path:        "/auth/signup",
			Middlewares: []echo.MiddlewareFunc{},
			Method:      http.MethodPost,
		},
		{
			Func:        userController.Index,
			Path:        "/users",
			Middlewares: []echo.MiddlewareFunc{},
			Method:      http.MethodGet,
		},
	}

	for _, route := range Routes {
		v1.Add(route.Method, route.Path, route.Func, route.Middlewares...)
	}

}
