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

type productID primitive.ObjectID

func (p *productID) UnmarshalJSON(data []byte) error {
	if len(string(data)) != 24 {
		(*p) = productID(primitive.NewObjectID())
		return nil
	}

	var err error
	var objID primitive.ObjectID

	if objID, err = primitive.ObjectIDFromHex(string(data)); err != nil {
		return err
	}

	(*p) = productID(primitive.ObjectID(objID))
	return nil
}

func (p *productID) MarshalJSON() ([]byte, error) {
	return primitive.ObjectID(*p).MarshalJSON()
}

// Product defines the model from
// OpenFoodFacts and the model that will be inserted at db
type Product struct {
	ID              productID     `json:"_id" bson:"_id,omitempty"`
	Code            string        `json:"code" bson:"code"`
	Status          ProductStatus `json:"status" bson:"status"`
	ImportedT       int64         `json:"imported_t" bson:"imported_t"`
	URL             string        `json:"url" bson:"url"`
	Creator         string        `json:"creator" bson:"creator"`
	CreatedT        int           `json:"created_t" bson:"created_t"`
	LastModifiedT   int           `json:"last_modified_t" bson:"last_modified_t"`
	ProductName     string        `json:"product_name" bson:"product_name"`
	Quantity        string        `json:"quantity" bson:"quantity"`
	Brands          string        `json:"brands" bson:"brands"`
	Categories      string        `json:"categories" bson:"categories"`
	Labels          string        `json:"labels" bson:"labels"`
	Cities          string        `json:"cities" bson:"cities"`
	PurchasePlaces  string        `json:"purchase_places" bson:"purchase_places"`
	Stores          string        `json:"stores" bson:"stores"`
	IngredientsText string        `json:"ingredients_text" bson:"ingredients_text"`
	Traces          string        `json:"traces" bson:"traces"`
	ServingSize     string        `json:"serving_size" bson:"serving_size"`
	ServingQuantity productFloat  `json:"serving_quantity" bson:"serving_quantity"`
	NutriscoreScore int           `json:"nutriscore_score" bson:"nutriscore_score"`
	NutriscoreGrage string        `json:"nutriscore_grade" bson:"nutriscore_grade"`
	MainCategory    string        `json:"main_category" bson:"main_category"`
	ImageURL        string        `json:"image_url" bson:"image_url"`
}
