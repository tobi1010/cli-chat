package cache

import (
	"cli-chat/fileatomic"
	"encoding/json"
	"fmt"
	"os"
)

func Load(cachePath string) (*Cache, error) {
	b, err := os.ReadFile(cachePath)
	if err != nil {
		return nil, fmt.Errorf("reading cache file %s: %w", cachePath, err)
	}

	var cache = Cache{}
	if err := json.Unmarshal(b, &cache); err != nil {
		return nil, fmt.Errorf("unmarshalling json: %w", err)
	}

	if cache.Providers == nil {
		cache.Providers = make(map[string]ModelList)
	}
	return &cache, nil
}

func (c *Cache) Save(cachePath string) error {
	b, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	if err := fileatomic.Write(cachePath, b, 0o600); err != nil {
		return fmt.Errorf("writing file atomically %s: %w", cachePath, err)
	}
	return nil
}
