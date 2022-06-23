package main

import (
	"fmt"
	"os"

	"github.com/ElioenaiFerrari/security-dog-api/entities"
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

	user := entities.User{
		UserName:   "Eli",
		Email:      "elioenaiferrari@gmail.com",
		Password:   "123456",
		Role:       "admin",
		Registries: []entities.Registry{},
	}

	db.Create(&user)

	db.Create(&entities.Registry{
		UserID:   user.ID,
		Name:     "registry1",
		Login:    "registry1",
		Password: "123456",
	})

	db.Preload("Registries").Find(&user)

	fmt.Println(user.Registries)

	server := echo.New()
	// v1 := server.Group("/api/v1")

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	server.Use(middleware.Gzip())
	server.Use(middleware.Secure())
	server.Use(middleware.CSRF())

	server.Logger.Fatal(server.Start(":4000"))
}
