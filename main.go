package main

import (
	"context"
	"log"
	"time"

	"go-challenge/config"
	"go-challenge/database"
	"go-challenge/internals/handlers"
	"go-challenge/internals/notification"
	"go-challenge/internals/repository"
	"go-challenge/internals/services"
	"go-challenge/pkg/httpclient"
	"go-challenge/server"

	"go.uber.org/fx"
)

// @title Go Open Food Facts Changelenge
// @version 1.0
// @description This project needs to realize synchronization with Open Food Facts open data and allow CRUD operations with data
// @termsOfService http://swagger.io/terms/
// @contact.name Eclésio F Melo Júnior
// @contact.url https://ecles.io
// @contact.email eclesiomelo.1@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host http://localhost:8080
// @BasePath /
func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
			database.NewConnection,
			httpclient.NewHTTPClient,
			notification.NewEmailer,
			notification.NewImportationNotifier,
		),

		fx.Provide(
			repository.NewProductRespository,
			repository.NewImportRepository,
		),

		fx.Provide(
			services.NewHealthcheck,
			services.NewImportation,
			services.NewProduct,

			handlers.NewHealthcheckHandler,
			handlers.NewProductsHandler,

			server.NewServer,
		),

		fx.Invoke(
			server.ImportJob,
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
