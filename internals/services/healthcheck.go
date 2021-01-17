package services

import (
	"context"
	"runtime"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Healthcheck interface define functions
// that returns the database connection status
// last time the sync was done and the system status
type Healthcheck interface {
	DatabaseReady() (bool, error)
	LastSyncExecution() time.Time
	GetMemUsage() uint64

	SetOnlineSince(time.Time)
	OnlineSince() time.Duration

	SetLastSync(time.Time)
}

type hc struct {
	lastSync    time.Time
	onlineSince time.Time
	mongo       *mongo.Database
}

// NewHealthcheck returns an implementation of Healthcheck interface
func NewHealthcheck(m *mongo.Database) Healthcheck {
	return &hc{
		mongo: m,
	}
}

func (h *hc) DatabaseReady() (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var err error
	var collections []string

	tmpCollection := "tmp_hc"
	tmpCollectionExists := false

	if collections, err = h.mongo.ListCollectionNames(ctx, bson.D{}); err != nil {
		return false, err
	}

	for _, coll := range collections {
		if coll == tmpCollection {
			tmpCollectionExists = true
		}
	}

	if !tmpCollectionExists {
		if err = h.mongo.CreateCollection(ctx, tmpCollection); err != nil {
			return false, err
		}
	}

	tmpDocument := struct {
		Value string `json:"value"`
	}{
		Value: "hc_insert_data",
	}

	if _, err = h.mongo.Collection(tmpCollection).InsertOne(ctx, tmpDocument); err != nil {
		return false, err
	}

	if err = h.mongo.Collection(tmpCollection).Drop(ctx); err != nil {
		return false, err
	}

	return true, nil
}

func (h *hc) LastSyncExecution() time.Time {
	return h.lastSync
}

func (h *hc) GetMemUsage() uint64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return m.Alloc
}

func (h *hc) SetOnlineSince(t time.Time) {
	h.onlineSince = t
}

func (h *hc) SetLastSync(t time.Time) {
	h.lastSync = t
}

func (h *hc) OnlineSince() time.Duration {
	return time.Since(h.onlineSince)
}
