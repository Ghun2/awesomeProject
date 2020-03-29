package api

import (
	"github.com/labstack/echo/v4"
	"github/Ghun2/awesomeProject/api/handlers"
)

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
