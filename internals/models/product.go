package models

import "time"

// ProductStatus enum the product status model
type ProductStatus string

// ProductStatus "enum" items
const (
	Draft     ProductStatus = "draft"
	Trash                   = "trash"
	Published               = "published"
)

// Product defines the model from
// OpenFoodFacts and the model that will be inserted at db
type Product struct {
	Code            int           `json:"code"`
	Status          ProductStatus `json:"status"`
	ImportedT       time.Time     `json:"imported_t"`
	URL             string        `json:"url"`
	Creator         string        `json:"creator"`
	CreatedT        time.Time     `json:"created_t"`
	LastModifiedT   time.Time     `json:"last_modified_t"`
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
	ServingQuantity float64       `json:"serving_quantity"`
	NutriscoreScore int           `json:"nutriscore_score"`
	NutriscoreGrage string        `json:"nutriscore_grade"`
	MainCategory    string        `json:"main_category"`
	ImageURL        string        `json:"image_url"`
}
