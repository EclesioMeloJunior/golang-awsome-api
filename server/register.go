package server

import (
	"go-challenge/internals/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	// Generate automatically the swagger docs
	_ "go-challenge/docs"
)

// NewRegister will setup the middlewares request endpoint handlers and inject the necessary deps
func NewRegister(e *echo.Echo, hc *handlers.Healthcheck, p *handlers.Products) {
	e.GET("/", hc.GetAPIStatus)

	e.GET("/products/:pcode", p.GetProductByID)
	e.PUT("/products/:pcode", p.UpdateProductByID)
	e.DELETE("/products/:pcode", p.RemoveProductByID)

	e.GET("/products", p.GetProductsList)

	e.GET("/swagger/*any", echoSwagger.WrapHandler)
}
