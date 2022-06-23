package main

import (
	"os"

	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"github.com/ElioenaiFerrari/security-dog-api/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func init() {
	godotenv.Load()
}

func main() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_DSN")), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Registry{})

	server := echo.New()
	v1 := server.Group("/api/v1")

	router.InitV1(v1, db)

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	server.Use(middleware.Gzip())
	server.Use(middleware.Secure())
	server.Use(middleware.CSRF())

	server.Logger.Fatal(server.Start(":4000"))
}
