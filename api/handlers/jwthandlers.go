package handlers

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func MainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)

	claims := token.Claims.(jwt.MapClaims)

	log.Print("User Name: ", claims["name"], " User ID: ", claims["jti"])

	return c.String(http.StatusOK, "you are on the top secret jwt page!")
}
