package main

import (
	"context"
	"log"
	"time"

	"go-challenge/config"
	"go-challenge/internals/handlers"
	"go-challenge/internals/services"
	"go-challenge/server"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			services.NewHealthcheck,
			handlers.NewHealthcheckHandler,
			server.NewServer,
		),

		fx.Invoke(server.NewRegister),
	)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}
