package main

import (
	"os"
	"time"

	"github.com/ElioenaiFerrari/security-dog-api/entities"
	"github.com/ElioenaiFerrari/security-dog-api/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/patrickmn/go-cache"
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
	db.AutoMigrate(&entities.Device{})

	server := echo.New()
	v1 := server.Group("/api/v1")

	memory := cache.New(time.Minute, time.Minute*15)

	router.InitV1(v1, memory, db)

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	server.Use(middleware.Gzip())
	server.Use(middleware.Secure())
	server.Use(middleware.BodyLimit("2M"))

	server.Logger.Fatal(server.Start(":4000"))
}
