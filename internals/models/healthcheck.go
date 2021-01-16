package models

import "time"

// Status struct define a simple fields
// to get the application status
type Status struct {
	OnlineT  string    `json:"online_t"`
	MemUsage string    `json:"mem_usage"`
	LastSync time.Time `json:"last_sync"`
	Database struct {
		Status      string `json:"status"`
		Description string `json:"description"`
	} `json:"database"`
}
