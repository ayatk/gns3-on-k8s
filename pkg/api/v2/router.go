package v2

import "github.com/labstack/echo/v4"

func Router(echo *echo.Echo) *echo.Group {
	g := echo.Group("/v2")
	g.GET("/version", VersionHandler)
	g.POST("/version", CheckVersionHandler)
	return g
}
