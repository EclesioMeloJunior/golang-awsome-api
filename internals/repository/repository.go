package repository

import (
	"context"
	"time"
)

func createContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 15*time.Second)
}
