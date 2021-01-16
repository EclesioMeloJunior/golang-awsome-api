package server

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// NewServer returns a pointer to Server
func NewServer(lc fx.Lifecycle) *echo.Echo {
	e := echo.New()

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go e.Start(":8080")
			return nil
		},
		OnStop: func(c context.Context) error {
			log.Println("Stopping server")
			return e.Shutdown(c)
		},
	})

	return e
}
