package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Import defines the model to interact with mongodb
// the field StopedAt helps to failed imports to start
// from a specific index inside de array
type Import struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Filename  string             `json:"filename" bson:"filename,omitempty"`
	ImportedT int64              `json:"imported_t" bson:"imported_t,omitempty"`
	Quantity  int                `json:"quantity" bson:"quantity,omitempty"`
}
