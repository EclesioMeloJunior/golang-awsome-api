package healthcheck

import "time"

// Status struct define a simple fields
// to get the application status
type Status struct {
	OnlineT  time.Time `json:"online_t"`
	MemUsage float64   `json:"mem_usage"`
	LastSync time.Time `json:"last_sync"`
	Database struct {
		Status      string `json:"status"`
		Description string `json:"description"`
	} `json:"database"`
}

// Healthcheck interface define functions
// that returns the database connection status
// last time the sync was done and the system status
type Healthcheck interface {
	DatabaseReady() (bool, error)
	LastSyncExecution() (time.Time, error)
	System() (time.Time, float64)
}

type hc struct{}

// NewHealthcheck returns an implementation of Healthcheck interface
func NewHealthcheck() Healthcheck {
	return &hc{}
}

func (h *hc) DatabaseReady() (bool, error) {
	return true, nil
}

func (h *hc) LastSyncExecution() (time.Time, error) {
	return time.Now(), nil
}

func (h *hc) System() (time.Time, float64) {
	return time.Now(), 0
}
