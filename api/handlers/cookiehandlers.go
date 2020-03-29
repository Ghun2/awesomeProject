package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret cookie page!")
}