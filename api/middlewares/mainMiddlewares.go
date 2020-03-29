package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetMainMiddlewares(e *echo.Echo) {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:"[${time_rfc3339}]  ${status}  ${method}  ${host}${path} ${latency_human}" + "\n",
	}))
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root: "static",
		HTML5: true,
	}))

	e.Use(serverHeader)
}


func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("notReallyHeader", "thisHaveNoMearning")
		return next(c)
	}
}
