package repository

import (
	"go-challenge/internals/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Product abstract the interactions between
// application and database
type Product interface {
	GetProducts(filter interface{}, page int, size int) ([]models.Product, error)
	InsertManyProducts([]interface{}, mongo.Session) error
}

// ProductsCollection is the name on mongodb
const ProductsCollection = "products"

type product struct {
	*mongo.Database
}

// NewProductRespository receives the mongo db instance to
// executes the operation
func NewProductRespository(m *mongo.Database) Product {
	return &product{m}
}

func (p *product) GetProducts(filter interface{}, page int, size int) ([]models.Product, error) {
	ctx, cancel := createContext()
	defer cancel()

	if filter == nil {
		filter = bson.D{}
	}

	findOpts := options.Find()

	findOpts.SetLimit(int64(size))
	findOpts.SetSkip(int64(size * page))

	var err error
	var c *mongo.Cursor

	if c, err = p.Database.Collection(ProductsCollection).Find(ctx, filter, findOpts); err != nil {
		return nil, err
	}

	products := make([]models.Product, 0)
	if err = c.All(ctx, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (p *product) InsertManyProducts(products []interface{}, session mongo.Session) error {
	ctx, cancel := createContext()
	defer cancel()

	var err error

	return mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		if _, err = p.Database.Collection(ProductsCollection).InsertMany(ctx, products); err != nil {
			return err
		}

		return nil
	})
}
