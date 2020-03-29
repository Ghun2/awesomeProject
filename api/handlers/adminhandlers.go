package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "ji hun you are on the secret admin main page!")
}