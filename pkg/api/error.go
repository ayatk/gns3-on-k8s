package api

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Error struct {
	StatusCode  int    `json:"status_code"`
	RespondedAt string `json:"responded_at"` // RFC3339
	Message     string `json:"message"`
}

func apiErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	message := ""
	// https://godoc.org/github.com/labstack/echo#HTTPError
	if ee, ok := err.(*echo.HTTPError); ok {
		code = ee.Code
		message = ee.Message.(string)
	}
	body := Error{
		StatusCode:  code,
		RespondedAt: time.Now().Format(time.RFC3339),
		Message:     message,
	}
	c.JSONPretty(code, body, "  ")
}
