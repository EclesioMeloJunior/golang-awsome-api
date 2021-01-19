package services

import (
	"go-challenge/internals/models"
	"go-challenge/internals/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Product interface handle the product business logic
// and interaction with repositories
type Product interface {
	GetProducts(filter interface{}, page int, size int) ([]models.Product, error)
	GetProductByID(string) (*models.Product, error)
	UpdateProductByID(string, *models.Product) (*models.Product, error)
	DeleteProductByID(string) (*models.Product, error)
}

type product struct {
	productRepo repository.Product
}

// NewProduct returns an implementaion of Products
func NewProduct(p repository.Product) Product {
	return &product{
		productRepo: p,
	}
}

func (p *product) GetProducts(filter interface{}, page int, size int) ([]models.Product, error) {
	return p.productRepo.GetProducts(filter, page, size)
}

func (p *product) GetProductByID(id string) (*models.Product, error) {
	var err error
	var objID primitive.ObjectID

	if objID, err = primitive.ObjectIDFromHex(id); err != nil {
		return p.productRepo.GetProductByCode(id)
	}

	return p.productRepo.GetProductByID(objID)
}

func (p *product) UpdateProductByID(id string, update *models.Product) (*models.Product, error) {
	var err error

	var objID primitive.ObjectID

	if objID, err = primitive.ObjectIDFromHex(id); err != nil {
		err = p.productRepo.UpdateProductByCode(id, update)
	} else {
		err = p.productRepo.UpdateProductByID(objID, update)
	}

	if err != nil {
		return nil, err
	}

	return p.GetProductByID(id)
}

func (p *product) DeleteProductByID(id string) (*models.Product, error) {
	var err error
	var product *models.Product

	if product, err = p.GetProductByID(id); err != nil {
		return nil, err
	}

	product.ToTrash()

	return p.UpdateProductByID(id, product)
}
