package models

import (
	"strconv"
	"strings"
)

// ProductStatus enum the product status model
type ProductStatus string

// ProductStatus "enum" items
const (
	Draft     ProductStatus = "draft"
	Trash                   = "trash"
	Published               = "published"
)

type productFloat int

func (p *productFloat) UnmarshalJSON(data []byte) error {
	f, err := strconv.ParseFloat(strings.Trim(string(data), "\"\""), 64)

	if err != nil {
		return err
	}

	(*p) = productFloat(f)
	return nil
}

// Product defines the model from
// OpenFoodFacts and the model that will be inserted at db
type Product struct {
	Code            string        `json:"code"`
	Status          ProductStatus `json:"status"`
	ImportedT       int           `json:"imported_t"`
	URL             string        `json:"url"`
	Creator         string        `json:"creator"`
	CreatedT        int           `json:"created_t"`
	LastModifiedT   int           `json:"last_modified_t"`
	ProductName     string        `json:"product_name"`
	Quantity        string        `json:"quantity"`
	Brands          string        `json:"brands"`
	Categories      string        `json:"categories"`
	Labels          string        `json:"labels"`
	Cities          string        `json:"cities"`
	PurchasePlaces  string        `json:"purchase_places"`
	Stores          string        `json:"stores"`
	IngredientsText string        `json:"ingredients_text"`
	Traces          string        `json:"traces"`
	ServingSize     string        `json:"serving_size"`
	ServingQuantity productFloat  `json:"serving_quantity"`
	NutriscoreScore int           `json:"nutriscore_score"`
	NutriscoreGrage string        `json:"nutriscore_grade"`
	MainCategory    string        `json:"main_category"`
	ImageURL        string        `json:"image_url"`
}
