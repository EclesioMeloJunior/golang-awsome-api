package repository

import (
	"go-challenge/internals/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Import repository interface will abstract
// the interations with database
type Import interface {
	GetLastImport() (*models.Import, error)
	GetAllImports() ([]models.Import, error)
	ExecuteImport(*models.Import, []interface{}) error
}

// ImportsCollection is the name on mongodb
const ImportsCollection = "imports"

type importation struct {
	mongo       *mongo.Database
	productRepo Product
}

// NewImportRepository receives the mongo db instance to
// executes the operation
func NewImportRepository(m *mongo.Database, p Product) Import {
	return &importation{
		mongo:       m,
		productRepo: p,
	}
}

func (i *importation) GetLastImport() (*models.Import, error) {
	ctx, cancel := createContext()
	defer cancel()

	findOpts := options.FindOne()
	findOpts.SetSort(bson.D{{"imported_t", -1}})

	var err error
	imp := new(models.Import)

	result := i.mongo.Collection(ImportsCollection).FindOne(ctx, bson.D{}, findOpts)
	if err = result.Decode(imp); err != nil {
		return nil, err
	}

	return imp, nil
}

func (i *importation) GetAllImports() ([]models.Import, error) {
	ctx, cancel := createContext()
	defer cancel()

	var err error
	var c *mongo.Cursor

	if c, err = i.mongo.Collection(ImportsCollection).Find(ctx, bson.D{}); err != nil {
		return nil, err
	}

	imports := make([]models.Import, 0)

	for c.Next(ctx) {
		i := models.Import{}

		if err = c.Decode(&i); err != nil {
			return nil, err
		}

		imports = append(imports, i)
	}

	return imports, nil
}

func (i *importation) ExecuteImport(imp *models.Import, products []interface{}) error {
	ctx, cancel := createContext()
	defer cancel()

	var err error
	var session mongo.Session

	log.Printf("Transaction to %s created\n", imp.Filename)

	if session, err = i.mongo.Client().StartSession(); err != nil {
		return err
	}

	if err = session.StartTransaction(); err != nil {
		return err
	}

	return mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
		log.Printf("Insert products to %s\n", imp.Filename)

		if err = i.productRepo.InsertManyProducts(products, session); err != nil {
			session.AbortTransaction(ctx)
			return err
		}

		log.Printf("Insert import to %s\n", imp.Filename)

		if _, err = i.mongo.Collection(ImportsCollection).InsertOne(ctx, imp); err != nil {
			session.AbortTransaction(ctx)
			return err
		}

		if err = session.CommitTransaction(sc); err != nil {
			return err
		}

		log.Printf("Commited to %s\n", imp.Filename)

		return nil
	})
}
