package services

import "time"

// Healthcheck interface define functions
// that returns the database connection status
// last time the sync was done and the system status
type Healthcheck interface {
	DatabaseReady() (bool, error)
	LastSyncExecution() (time.Time, error)
	GetMemUsage() float64

	SetOnlineSince(time.Time)
	OnlineSince() time.Duration
}

type hc struct {
	onlineSince time.Time
}

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

func (h *hc) GetMemUsage() float64 {
	return 0
}

func (h *hc) SetOnlineSince(t time.Time) {
	h.onlineSince = t
}

func (h *hc) OnlineSince() time.Duration {
	return time.Since(h.onlineSince)
}
