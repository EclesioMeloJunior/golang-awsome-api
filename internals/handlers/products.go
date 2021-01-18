package handlers

import (
	"go-challenge/internals/models"
	"go-challenge/internals/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Products is a struct that has the product service dependency
type Products struct {
	productService services.Product
}

// NewProductsHandler will return a pointer to Product
// at construction will inject te product service dependency
func NewProductsHandler(p services.Product) *Products {
	return &Products{
		productService: p,
	}
}

// GetProductsList will return a list with products
func (p *Products) GetProductsList(c echo.Context) error {
	page, size := getPagination(c)

	var err error
	var products []models.Product

	if products, err = p.productService.GetProducts(nil, page, size); err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	return c.JSON(http.StatusOK, successResponse(products))
}
