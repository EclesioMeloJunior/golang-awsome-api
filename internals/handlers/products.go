package handlers

import (
	"errors"
	"go-challenge/internals/models"
	"go-challenge/internals/services"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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
// @Summary Products List
// @Description return a list of products imported from Open Food Facts open database
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{body=[]models.Product,success=bool}
// @Failure 500 {object} Response{message=string,success=bool}
// @Router /products [get]
func (p *Products) GetProductsList(c echo.Context) error {
	page, size := getPagination(c)

	var err error
	var products []models.Product

	products, err = p.productService.GetProducts(nil, page, size)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	return c.JSON(http.StatusOK, successResponse(products))
}

// GetProductByID will find a product by ID or code
// @Summary Get Product By ID or Code
// @Description return one product found by either ID or Code
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{body=models.Product,success=bool}
// @Success 204
// @Failure 500 {object} Response{message=string,success=bool}
// @Router /products/:idOrCode [get]
func (p *Products) GetProductByID(c echo.Context) error {
	productID := c.Param("pcode")

	var err error
	var product *models.Product

	product, err = p.productService.GetProductByID(productID)

	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNoContent, nil)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	return c.JSON(http.StatusOK, successResponse(product))
}

// UpdateProductByID will find and update the product by its ID
// @Summary Update One Product
// @Description Retrive a product by either ID or code and update info
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{body=models.Product,success=bool}
// @Failure 400 {object} Response{message=string,success=bool} "Validation fails"
// @Success 204
// @Failure 500 {object} Response{message=string,success=bool}
// @Router /products/:idOrCode [put]
func (p *Products) UpdateProductByID(c echo.Context) error {
	productID := c.Param("pcode")

	product := new(models.Product)

	var err error
	if err = c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	if product.ID != primitive.NilObjectID {
		err = errors.New("Invalid body data: _id")
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	if product.Code != "" {
		err = errors.New("Invalid body data: code")
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	if product.ImportedT != 0 {
		err = errors.New("Invalid body data: imported_t")
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	if product.Status != "" {
		err = errors.New("Invalid body data: status")
		return c.JSON(http.StatusBadRequest, errorResponse(err))
	}

	product, err = p.productService.UpdateProductByID(productID, product)

	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNoContent, nil)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	return c.JSON(http.StatusOK, successResponse(product))
}

// RemoveProductByID will find and update the product status to "trash"
// @Summary Remove One Product
// @Description Retrive a product by either ID or code and update the status to "trash"
// @Accept  json
// @Produce  json
// @Success 200 {object} Response{body=models.Product,success=bool}
// @Success 204
// @Failure 500 {object} Response{message=string,success=bool}
// @Router /products/:idOrCode [delete]
func (p *Products) RemoveProductByID(c echo.Context) error {
	productID := c.Param("pcode")

	var err error
	var product *models.Product

	product, err = p.productService.DeleteProductByID(productID)

	if err == mongo.ErrNoDocuments {
		return c.JSON(http.StatusNoContent, nil)
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	return c.JSON(http.StatusOK, successResponse(product))
}
