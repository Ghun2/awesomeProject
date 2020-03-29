package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hel412lo, Wo123rld!")
}