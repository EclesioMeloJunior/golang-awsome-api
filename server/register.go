package server

import (
	"go-challenge/internals/handlers"

	"github.com/labstack/echo/v4"
)

// NewRegister will setup the middlewares
// request endpoint handlers and inject
// the necessary dependecies
func NewRegister(e *echo.Echo, hc *handlers.Healthcheck, p *handlers.Products) {
	e.GET("/", hc.GetAPIStatus)
}
