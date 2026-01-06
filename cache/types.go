package cache

import (
	"cli-chat/internal/apitypes"
	"time"
)

type ModelList struct {
	FetchedAt time.Time        `json:"fetched_at"`
	Models    []apitypes.Model `json:"models"`
}

type Cache struct {
	Providers map[string]ModelList `json:"providers"`
}

func New() *Cache {
	return &Cache{
		Providers: make(map[string]ModelList),
	}
}
