package controllers

import (
	"context"

	"github.com/Lincoln-a-Bot/GoPost/web/templates"
	"github.com/labstack/echo/v4"
)

func LoginPage(c echo.Context) error {
	// Retrieve the error message from the query parameter
	errorMessage := c.QueryParam("error")
	// Pass the error message to the template
	component := templates.LoginPage(errorMessage)
	return component.Render(context.Background(), c.Response().Writer)
}
