package models

import (
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Code            string             `json:"code,omitempty" bson:"code,omitempty"`
	Status          ProductStatus      `json:"status,omitempty" bson:"status,omitempty"`
	ImportedT       int64              `json:"imported_t,omitempty" bson:"imported_t,omitempty"`
	URL             string             `json:"url,omitempty" bson:"url,omitempty"`
	Creator         string             `json:"creator,omitempty" bson:"creator,omitempty"`
	CreatedT        int                `json:"created_t,omitempty" bson:"created_t,omitempty"`
	LastModifiedT   int                `json:"last_modified_t,omitempty" bson:"last_modified_t,omitempty"`
	ProductName     string             `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Quantity        string             `json:"quantity,omitempty" bson:"quantity,omitempty"`
	Brands          string             `json:"brands,omitempty" bson:"brands,omitempty"`
	Categories      string             `json:"categories,omitempty" bson:"categories,omitempty"`
	Labels          string             `json:"labels,omitempty" bson:"labels,omitempty"`
	Cities          string             `json:"cities,omitempty" bson:"cities,omitempty"`
	PurchasePlaces  string             `json:"purchase_places,omitempty" bson:"purchase_places,omitempty"`
	Stores          string             `json:"stores,omitempty" bson:"stores,omitempty"`
	IngredientsText string             `json:"ingredients_text,omitempty" bson:"ingredients_text,omitempty"`
	Traces          string             `json:"traces,omitempty" bson:"traces,omitempty"`
	ServingSize     string             `json:"serving_size,omitempty" bson:"serving_size,omitempty"`
	ServingQuantity productFloat       `json:"serving_quantity,omitempty" bson:"serving_quantity,omitempty"`
	NutriscoreScore int                `json:"nutriscore_score,omitempty" bson:"nutriscore_score,omitempty"`
	NutriscoreGrage string             `json:"nutriscore_grade,omitempty" bson:"nutriscore_grade,omitempty"`
	MainCategory    string             `json:"main_category,omitempty" bson:"main_category,omitempty"`
	ImageURL        string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
}

// ToTrash updates the product status to "trash"
func (p *Product) ToTrash() {
	p.Status = Trash
}
