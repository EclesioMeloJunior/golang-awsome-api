package models

// Import defines the model to interact with mongodb
// the field StopedAt helps to failed imports to start
// from a specific index inside de array
type Import struct {
	Filename string `json:"filename"`
	Imported bool   `json:"imported"`
	StopedAt int    `json:"stoped_at"`
}
