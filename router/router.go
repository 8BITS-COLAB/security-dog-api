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

	auth := middleware.JWTWithConfig(config)
	// csrf := middleware.CSRF()

	authController := factories.MakeAuthController(db)
	userController := factories.MakeUserController(db)
	registryController := factories.MakeRegistryController(db)
	deviceController := factories.MakeDeviceController(db)
	sharedRegistryController := factories.MakeSharedRegistryController(db)
	detectiveController := factories.MakeDetectiveController()

	var routes = []Route{
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
			Func:        authController.Profile,
			Path:        "/auth/profile",
			Middlewares: []echo.MiddlewareFunc{auth},
			Method:      http.MethodGet,
		},
		{
			Func:        authController.RefreshToken,
			Path:        "/auth/refresh-token",
			Middlewares: []echo.MiddlewareFunc{auth},
			Method:      http.MethodGet,
		},
		{
			Func:        authController.TwoFA,
			Path:        "/auth/2fa",
			Middlewares: []echo.MiddlewareFunc{auth},
			Method:      http.MethodGet,
		},
		{
			Func: authController.CSRFToken,
			Path: "/auth/csrf-token",
			Middlewares: []echo.MiddlewareFunc{
				auth,
				// csrf,
			},
			Method: http.MethodGet,
		},
		// Users
		{
			Func: userController.Index,
			Path: "/users",
			Middlewares: []echo.MiddlewareFunc{
				auth,
			},
			Method: http.MethodGet,
		},
		{
			Func: userController.Show,
			Path: "/users/:id",
			Middlewares: []echo.MiddlewareFunc{
				auth,
			},
			Method: http.MethodGet,
		},
		{
			Func: userController.Update,
			Path: "/users/:id",
			Middlewares: []echo.MiddlewareFunc{
				auth,
				// csrf,
			},
			Method: http.MethodPatch,
		},
		{
			Func: userController.Delete,
			Path: "/users/:id",
			Middlewares: []echo.MiddlewareFunc{
				auth,
				// csrf,
			},
			Method: http.MethodDelete,
		},
		// Registries
		{
			Func: registryController.Index,
			Path: "/users/:user_id/registries",
			Middlewares: []echo.MiddlewareFunc{
				auth,
			},
			Method: http.MethodGet,
		},
		{
			Func: registryController.Create,
			Path: "/users/:user_id/registries",
			Middlewares: []echo.MiddlewareFunc{
				auth,
				// csrf,
			},
			Method: http.MethodPost,
		},
		{
			Func: registryController.Show,
			Path: "/users/:user_id/registries/:id",
			Middlewares: []echo.MiddlewareFunc{
				auth,
			},
			Method: http.MethodGet,
		},
		{
			Func: registryController.Update,
			Path: "/users/:user_id/registries/:id",
			Middlewares: []echo.MiddlewareFunc{
				auth,
				// csrf,
			},
			Method: http.MethodPatch,
		},
		{
			Func: registryController.Delete,
			Path: "/users/:user_id/registries/:id",
			Middlewares: []echo.MiddlewareFunc{
				auth,
				// csrf,
			},
			Method: http.MethodDelete,
		},
		// Devices
		{
			Func: deviceController.Index,
			Path: "/users/:user_id/devices",
			Middlewares: []echo.MiddlewareFunc{
				auth,
			},
			Method: http.MethodGet,
		},
		{
			Func: deviceController.Update,
			Path: "/users/:user_id/devices",
			Middlewares: []echo.MiddlewareFunc{
				auth,
			},
			Method: http.MethodPatch,
		},
		// Shared registries
		{
			Func:        sharedRegistryController.Show,
			Path:        "/shared-registries/:id",
			Middlewares: []echo.MiddlewareFunc{},
			Method:      http.MethodGet,
		},
		{
			Func: sharedRegistryController.Create,
			Path: "/users/:user_id/shared-registries",
			Middlewares: []echo.MiddlewareFunc{
				auth,
				// csrf,
			},
			Method: http.MethodPost,
		},
		// Detective
		{
			Func: detectiveController.Investigate,
			Path: "/detective",
			Middlewares: []echo.MiddlewareFunc{
				auth,
			},
			Method: http.MethodGet,
		},
	}

	for _, route := range routes {
		v1.Add(route.Method, route.Path, route.Func, route.Middlewares...)
	}

}
