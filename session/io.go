package session

import (
	"cli-chat/cache"
	"cli-chat/fileatomic"
	"cli-chat/paths"
	"cli-chat/providers"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

func Open(paths paths.Paths) (*Session, error) {
	s, err := NewDefaultSession()

	if err != nil {
		return nil, fmt.Errorf("creating default session: %w", err)
	}

	if err := applySettings(s, paths); err != nil {
		return nil, fmt.Errorf("applying settings to session: %w", err)
	}

	if err := applyDb(s, paths); err != nil {
		return nil, fmt.Errorf("applying db to session: %w", err)
	}

	lastChat, err := resolveChat(*s.DB, paths)
	if err != nil {
		return nil, fmt.Errorf("resloving last chat: %w", err)
	}
	s.Chat = lastChat

	state, ok, err := loadState(paths.StatePath)
	if err != nil {
		return nil, fmt.Errorf("loading state: %w", err)
	}
	if !ok {
		s.ProviderName = providers.Default
	} else {
		s.ProviderName = state.LastProvider
	}
	requestedModelID := ""
	if ok {
		requestedModelID = state.LastModelID
	}

	cch, err := cache.Load(paths.CachePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			cch = cache.New()
			if err := cch.Save(s.Paths.CachePath); err != nil {
				return nil, fmt.Errorf("saving cache file: %w", err)
			}

		} else {
			return nil, fmt.Errorf("loading cache file: %w", err)
		}
	}
	s.Cache = cch

	_, ok = s.Cache.Providers[s.ProviderName]
	if !ok {
		s.Cache.Add(s.ProviderName)
	}
	if err := s.Cache.EnsureFresh(context.Background(), paths.CachePath, s.ProviderName, time.Duration(s.AppSettings.TTL)*time.Second, s); err != nil {
		return nil, fmt.Errorf("refreshing cache: %w", err)
	}

	lastModel, ok := s.Cache.Get(s.ProviderName, requestedModelID)
	if !ok {
		def, ok := providers.Get(s.ProviderName)
		if !ok {
			return nil, fmt.Errorf("resolving provider %q", s.ProviderName)
		}
		lastModel, ok = s.Cache.Get(s.ProviderName, def.DefaultModel)
		if !ok {
			return nil, fmt.Errorf("resolving default model %s for provider %s", def.DefaultModel, s.ProviderName)
		}
	}
	s.ModelID = lastModel.ID
	s.ModelLabel = lastModel.Name

	return s, nil

}

func (s *Session) Save(sessionPath string) error {
	st := State{
		LastProvider: s.ProviderName, LastModelID: s.ModelID,
	}
	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	if err = fileatomic.Write(sessionPath, data, 0o600); err != nil {
		return fmt.Errorf("writing file atomically %s: %w", sessionPath, err)
	}

	return nil
}
