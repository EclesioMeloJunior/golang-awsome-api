package handlers_test

import (
	"go-challenge/internals/handlers"
	"go-challenge/mocks/internals/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var productServiceMock *services.Product

func createProductHandler() *handlers.Products {
	productServiceMock = &services.Product{}
	return handlers.NewProductsHandler(productServiceMock)
}

func TestGetProductsList(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rr := httptest.NewRecorder()

	h := createProductHandler()

	c := e.NewContext(req, rr)

	err := h.GetProductsList(c)

	assert.Nil(t, err)
	assert.
}
