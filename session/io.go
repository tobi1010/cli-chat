package session

import (
	"cli-chat/cache"
	"cli-chat/chat"
	"cli-chat/fileatomic"
	"cli-chat/index"
	"cli-chat/paths"
	"cli-chat/settings"
	"context"
	"encoding/json"
	"fmt"
	"time"
)

func Open(paths paths.Paths) (*Session, error) {
	s, err := NewDefaultSession()
	if err != nil {
		return nil, fmt.Errorf("creating default session: %w", err)
	}

	if err := applyDb(s, paths); err != nil {
		return nil, fmt.Errorf("applying db to session: %w", err)
	}

	if err := applySettings(s, paths); err != nil {
		return nil, fmt.Errorf("applying settings to session: %w", err)
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
	}
	s.ProviderName = state.LastProvider

	cache, err := cache.Load(paths.CachePath)
	if err != nil {
		return nil, fmt.Errorf("loading cache file: %w", err)
	}
	s.Cache = cache
	_, ok = cache.Providers[s.ProviderName]
	if !ok {
		s.Cache.Add(s.ProviderName)
	}
	c, err := cache.EnsureFresh(context.Background(), paths.CachePath, s.ProviderName, time.Duration(s.AppSettings.TTL)*time.Second, s)

	return nil, nil

}

func Load(sessionPath string) (*Session, error) {
	s, err := NewDefaultSession()
	if err != nil {
		return nil, fmt.Errorf("creating default session: %w", err)
	}

	set, err := settings.LoadOrCreate()
	if err != nil {
		return nil, fmt.Errorf("load or create settings: %w", err)
	}
	s.applySettings(set)

	state, ok, err := loadState(sessionPath)
	if err != nil {
		return nil, fmt.Errorf("loading session file: %w", err)
	}
	if ok {
		s.applySavedState(state)
	}

	if err := s.loadDb(); err != nil {
		return nil, fmt.Errorf("loading database: %w", err)
	}

	lastChatId := s.DB.GetLastChatId()
	if err := s.loadChat(lastChatId); err != nil {
		return nil, fmt.Errorf("loading chat: %w", err)
	}

	return s, nil
}

func (s *Session) Save(sessionPath string) error {
	st := s.toState()
	data, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	if err = fileatomic.Write(sessionPath, data, 0o600); err != nil {
		return fmt.Errorf("writing file atomically %s: %w", sessionPath, err)
	}

	return nil
}

func applyDb(s *Session, paths paths.Paths) error {
	db, err := index.Load(paths.IndexPath)
	if err != nil {
		//create new db
		db = index.NewDB()
		s.DB = db
		err := db.Save(paths.IndexPath)
		if err != nil {
			return fmt.Errorf("saving db file: %w", err)
		}
	}
	return nil
}
func applySettings(s *Session, paths paths.Paths) error {
	set, err := settings.Load(paths.SettingsPath)
	if err != nil {
		//apply default settings
		set = settings.NewDefaultSettings()
		s.AppSettings = set
		err := s.AppSettings.Save(paths.SettingsPath)
		if err != nil {
			return fmt.Errorf("saving settings file: %w", err)
		}
	}
	s.AppSettings = set
	return nil
}
func resolveChat(db index.DB, paths paths.Paths) (*chat.Chat, error) {
	lastChatID := db.GetLastChatId()
	if lastChatID == "" {
		return chat.New(), nil
	} else {
		c, err := chat.Load(paths.ChatsDir, lastChatID)
		if err != nil {
			return nil, fmt.Errorf("readign last chat, id: %s: %w", lastChatID, err)
		}
		return c, nil
	}
}
