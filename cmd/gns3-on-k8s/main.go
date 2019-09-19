package main

import (
	"github.com/ayatk/gns3-on-k8s/pkg/api"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	api.Router(e)

	// Start server
	e.Logger.Fatal(e.Start(":5678"))
}
