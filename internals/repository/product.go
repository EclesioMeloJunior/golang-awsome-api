package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Product interface {
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
