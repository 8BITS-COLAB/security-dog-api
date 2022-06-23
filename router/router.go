package router

import (
	"net/http"
	"os"

	"github.com/ElioenaiFerrari/security-dog-api/factories"
	"github.com/ElioenaiFerrari/security-dog-api/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

type Route struct {
	Func        echo.HandlerFunc
	Path        string
	Middlewares []echo.MiddlewareFunc
	Method      string
}

func InitV1(v1 *echo.Group, db *gorm.DB) {

	var config = middleware.JWTConfig{
		Claims:     &security.JwtClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	authController := factories.MakeAuthController(db)
	userController := factories.MakeUserController(db)
	registryController := factories.MakeRegistryController(db)
	deviceController := factories.MakeDeviceController(db)

	var Routes = []Route{
		// Auth
		{
			Func:        authController.Signup,
			Path:        "/auth/signup",
			Middlewares: []echo.MiddlewareFunc{},
			Method:      http.MethodPost,
		},
		{
			Func:        authController.Signin,
			Path:        "/auth/signin",
			Middlewares: []echo.MiddlewareFunc{},
			Method:      http.MethodPost,
		},
		{
			Func:        authController.Signout,
			Path:        "/auth/signout",
			Middlewares: []echo.MiddlewareFunc{middleware.JWTWithConfig(config)},
			Method:      http.MethodGet,
		},
		{
			Func:        authController.Profile,
			Path:        "/auth/profile",
			Middlewares: []echo.MiddlewareFunc{middleware.JWTWithConfig(config)},
			Method:      http.MethodGet,
		},
		// Users
		{
			Func: userController.Index,
			Path: "/users",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWTWithConfig(config),
			},
			Method: http.MethodGet,
		},
		{
			Func: userController.Show,
			Path: "/users/:id",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWTWithConfig(config),
			},
			Method: http.MethodGet,
		},
		{
			Func: userController.Update,
			Path: "/users/:id",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWTWithConfig(config),
			},
			Method: http.MethodPatch,
		},
		{
			Func: userController.Delete,
			Path: "/users/:id",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWTWithConfig(config),
			},
			Method: http.MethodDelete,
		},
		// Registries
		{
			Func: registryController.Index,
			Path: "/users/:user_id/registries",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWTWithConfig(config),
			},
			Method: http.MethodGet,
		},
		{
			Func: registryController.Create,
			Path: "/users/:user_id/registries",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWTWithConfig(config),
			},
			Method: http.MethodPost,
		},
		{
			Func: registryController.Show,
			Path: "/users/:user_id/registries/:id",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWTWithConfig(config),
			},
			Method: http.MethodGet,
		},
		{
			Func: registryController.Update,
			Path: "/users/:user_id/registries/:id",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWTWithConfig(config),
			},
			Method: http.MethodPatch,
		},
		{
			Func: registryController.Delete,
			Path: "/users/:user_id/registries/:id",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWTWithConfig(config),
			},
			Method: http.MethodDelete,
		},
		// Devices
		{
			Func: deviceController.Index,
			Path: "/users/:user_id/devices",
			Middlewares: []echo.MiddlewareFunc{
				middleware.JWTWithConfig(config),
			},
			Method: http.MethodGet,
		},
	}

	for _, route := range Routes {
		v1.Add(route.Method, route.Path, route.Func, route.Middlewares...)
	}

}
