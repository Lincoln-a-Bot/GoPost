package main

import (
	"fmt"
	"net/http"

	"github.com/Lincoln-a-Bot/GoPost/config"
	"github.com/Lincoln-a-Bot/GoPost/internal/controllers"
	"github.com/Lincoln-a-Bot/GoPost/internal/models"
	"github.com/Lincoln-a-Bot/GoPost/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Creating echo server
	server := echo.New()

	server.Use(middleware.Logger())

	// Creating DB or connecting
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}

	// Creating user
	if err := models.CreateUser(db, "admin", "admin", "admin@example.com"); err != nil {
		fmt.Println(err)
	}

	// 404 handler
	server.HTTPErrorHandler = func(err error, c echo.Context) {
		if httpErr, ok := err.(*echo.HTTPError); ok {
			if httpErr.Code == http.StatusNotFound {
				c.String(http.StatusOK, "404 Page Not Found")
			}
		}
	}

	// Routing

	server.GET("/", controllers.HomePage)

	server.GET("/login", controllers.LoginPage)

	server.POST("/submit-login", func(c echo.Context) error {
		Username := c.FormValue("Username")
		Password := c.FormValue("Password")

		// Validate the credentials
		CheckLogin := services.CheckLogin(db, Username, Password)
		if CheckLogin {
			// Set the HX-Redirect header to redirect the client
			c.Response().Header().Set("HX-Redirect", "/")
			return nil
		}
		//htmx only accepts status code 200, so returning 200 and not a 401 StatusUnauthorized
		return c.HTML(http.StatusOK, "Invalid login")

	})

	// Starting server
	server.Logger.Fatal(server.Start(":8080"))
}
