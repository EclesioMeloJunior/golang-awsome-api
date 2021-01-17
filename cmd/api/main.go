package main

import (
	"context"
	"log"
	"time"

	"go-challenge/config"
	"go-challenge/database"
	"go-challenge/internals/handlers"
	"go-challenge/internals/repository"
	"go-challenge/internals/services"
	"go-challenge/pkg/httpclient"
	"go-challenge/server"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			database.NewConnection,
			httpclient.NewHTTPClient,
		),

		fx.Provide(
			repository.NewImportation,
		),

		fx.Provide(
			services.NewHealthcheck,
			services.NewImportation,
			handlers.NewHealthcheckHandler,
			handlers.NewProductsHandler,
			server.NewServer,
		),

		fx.Invoke(
			server.NewRegister,
		),
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
