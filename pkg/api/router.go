package api

import (
	v2 "github.com/ayatk/gns3-on-k8s/pkg/api/v2"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo) {
	e.HTTPErrorHandler = apiErrorHandler
	v2.Router(e)
}
