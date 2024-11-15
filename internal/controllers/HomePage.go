package controllers

import (
	"context"

	"github.com/Lincoln-a-Bot/GoPost/web/templates"
	"github.com/labstack/echo/v4"
)

func HomePage(c echo.Context) error {
	Component := templates.HomePage()
	return Component.Render(context.Background(), c.Response().Writer)
}
