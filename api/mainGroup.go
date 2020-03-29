package api

import (
	"github.com/labstack/echo/v4"
	"net/http"

	"github/Ghun2/awesomeProject/api/handlers"
)

func MainGroup(e *echo.Echo) {
	e.GET("/hello", handlers.Hello)
	e.GET("/login", handlers.Login)

	e.GET("/cats/:data", handlers.GetCats)

	e.GET("/cats", func (c echo.Context) error {
		return c.String(http.StatusOK, "Hel412lo, Wo123rld!")
	})

	e.POST("/cats", handlers.AddCat)
	e.POST("/dogs", handlers.AddDog)
	e.POST("/hamsters", handlers.AddHamster)
}
