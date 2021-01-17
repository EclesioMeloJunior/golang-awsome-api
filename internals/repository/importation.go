package repository

import (
	"go-challenge/internals/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Importation repository interface will abstract
// the interations with database
type Importation interface {
	GetAllImports() ([]models.Import, error)
}

// ImportsCollection is the name on mongodb
const ImportsCollection = "imports"

type importation struct {
	*mongo.Database
}

// NewImportation receives the mongo db instance to
// executes the operation
func NewImportation(m *mongo.Database) Importation {
	return &importation{m}
}

func (i *importation) GetAllImports() ([]models.Import, error) {
	ctx, cancel := createContext()
	defer cancel()

	var err error
	var c *mongo.Cursor

	if c, err = i.Database.Collection(ImportsCollection).Find(ctx, bson.D{}); err != nil {
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
