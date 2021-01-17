package handlers

import (
	"go-challenge/internals/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Products ...
type Products struct {
	importationService services.Importation
}

// NewProductsHandler ...
func NewProductsHandler(importationService services.Importation) *Products {
	return &Products{importationService}
}

// Import ...
func (p *Products) Import(e echo.Context) error {
	filenames, err := p.importationService.GetFilenames()

	if err != nil {
		return e.JSON(http.StatusInternalServerError, err)
	}

	imports, err := p.importationService.ToBeImported(filenames)

	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	product, err := p.importationService.ImportFiles(imports)

	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}

	return e.JSON(http.StatusOK, product)
}
