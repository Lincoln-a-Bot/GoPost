package main

import (
	"fmt"

	"github.com/Lincoln-a-Bot/GoPost/config"
	"github.com/Lincoln-a-Bot/GoPost/internal/controllers"
	"github.com/Lincoln-a-Bot/GoPost/internal/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	server := echo.New()

	server.Use(middleware.Logger())

	server.Static("/static", "web/static")

	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Println(err)
	}

	if err := models.CreateUser(db, "admin", "admin", "admin@example.com"); err != nil {
		fmt.Println(err)
	}

	// Register the renderer

	server.GET("/", controllers.HomePage)

	server.Logger.Fatal(server.Start(":8080"))
}
