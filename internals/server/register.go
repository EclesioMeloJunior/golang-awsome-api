package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewRegister will setup the middlewares
// request endpoint handlers and inject
// the necessary dependecies
func NewRegister(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

}
