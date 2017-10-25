package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HomeIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "home", nil)
}
