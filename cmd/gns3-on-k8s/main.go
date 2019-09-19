package main

import (
	v2 "github.com/ayatk/gns3-on-k8s/pkg/api/v2"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	v2.Router(e)

	// Start server
	e.Logger.Fatal(e.Start(":5678"))
}
