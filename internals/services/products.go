package services

import (
	"go-challenge/internals/models"
	"go-challenge/internals/repository"
)

// Product interface handle the product business logic
// and interaction with repositories
type Product interface {
	GetProducts(filter interface{}, page int, size int) ([]models.Product, error)
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
