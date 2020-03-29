package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetAdminMiddlewares(g *echo.Group) {
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// check in the DB
		if username == "hun" && password == "3133" {
			return true, nil;
		}
		return false, nil;
	}))
}
