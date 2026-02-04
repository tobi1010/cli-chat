package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"terminal-chat/cache"
	"terminal-chat/debug"
	"terminal-chat/fileatomic"
	"terminal-chat/paths"
	"terminal-chat/providers"
	"time"
)

func SessionOpen(paths paths.Paths) (*Session, error) {
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
	s.ProviderName = providers.Default

	if s.AppSettings.DefaultProvider != "" {
		s.ProviderName = s.AppSettings.DefaultProvider
	}

	requestedModelID := ""
	if s.AppSettings.DefaultModel != "" {
		requestedModelID = s.AppSettings.DefaultModel
	}

	state, ok, err := loadState(paths.SessionPath)
	if err != nil {
		return nil, fmt.Errorf("loading state: %w", err)
	}
	if ok {
		if _, exists := providers.Get(state.LastProvider); exists {
			s.ProviderName = state.LastProvider
		}
	}
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
	debug.Dump("Session:", s)

	return s, nil

}

func (s *Session) SessionSave(sessionPath string) error {
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
