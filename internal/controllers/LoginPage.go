package controllers

import (
	"context"

	"github.com/Lincoln-a-Bot/GoPost/web/templates"
	"github.com/labstack/echo/v4"
)

func LoginPage(c echo.Context) error {

	// Pass the error message to the template
	component := templates.LoginPage()
	return component.Render(context.Background(), c.Response().Writer)
}
