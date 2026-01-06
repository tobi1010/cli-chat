package cache

import (
	"cli-chat/internal/apitypes"
	"context"
	"fmt"
	"time"
)

func (c *Cache) Get(provider string, model string) (apitypes.Model, bool) {
	models, ok := c.Providers[provider]
	if !ok {
		return apitypes.Model{}, false
	}
	m, ok := models.getById(model)
	if !ok {
		return apitypes.Model{}, false
	}
	return m, true

}
func (m ModelList) getById(id string) (apitypes.Model, bool) {
	for _, model := range m.Models {
		if model.ID == id {
			return model, true
		}
	}
	return apitypes.Model{}, false
}

type ModelFetcher interface {
	FetchModels(ctx context.Context, provider string) ([]apitypes.Model, error)
}

func (c *Cache) EnsureFresh(ctx context.Context, cachePath string, provider string, ttl time.Duration, fetcher ModelFetcher) error {
	if c.Providers == nil {
		c.Providers = make(map[string]ModelList)
	}

	list := c.Providers[provider]
	if list.FetchedAt.IsZero() || list.IsStale(ttl) {
		models, err := fetcher.FetchModels(ctx, provider)
		if err != nil {
			return fmt.Errorf("fetching models for %q: %w", provider, err)
		}
		list.FetchedAt = time.Now()
		list.Models = models
		c.Providers[provider] = list
	}
	if err := c.Save(cachePath); err != nil {
		return fmt.Errorf("writing cache atomically: %w", err)
	}
	return nil
}

func (m ModelList) IsStale(ttl time.Duration) bool {
	return time.Since(m.FetchedAt) > ttl
}
