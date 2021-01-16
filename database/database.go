package database

import (
	"context"
	"go-challenge/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

// NewConnection will manage when the database connects
// and will stop connection when application shutdown
func NewConnection(lc fx.Lifecycle, c *config.Config) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(c.ConnStr))

	if err != nil {
		log.Fatal(err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Println("Connecting with database")
			return client.Connect(ctx)
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Disconnection from database")
			return client.Disconnect(ctx)
		},
	})

	return client
}
