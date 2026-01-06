package session

import (
	"cli-chat/cache"
	"cli-chat/chat"
	"cli-chat/index"
	"cli-chat/internal/client"
	"cli-chat/paths"
	"cli-chat/providers"
	"cli-chat/settings"
	"fmt"
	"time"
)

type State struct {
	LastProvider string `json:"last_provider"`
	LastModelID  string `json:"last_model_id"`
}
type Paths struct {
	SessionPath  string `json:"session_path"`
	ChatsDir     string `json:"chats_dir"`
	IndexPath    string `json:"index_path"`
	SettingsPath string `json:"settings_path"`
	CachePath    string `json:"cahce_path"`
}

type Session struct {
	Provider    providers.Provider
	Chat        *chat.Chat
	Client      *client.Client
	AppSettings settings.Settings
	DB          *index.DB
	Paths       Paths
	Cache       *cache.Cache
}

func NewDefaultSession() (*Session, error) {
	s := Session{}
	s.AppSettings = settings.NewDefaultSettings()
	s.Client = client.New(time.Duration(s.AppSettings.Timeout) * time.Second)
	s.Chat = chat.New()
	s.DB = index.NewDB()
	s.Cache = cache.New()
	provider, err := providers.NewDefault(*s.Cache)
	if err != nil {
		return nil, fmt.Errorf("default provider: %w", err)
	}
	s.Provider = provider

	sessionPath, err := paths.SessionPath()
	if err != nil {
		return nil, fmt.Errorf("resolving session path: %w", err)
	}
	chatsDir, err := paths.ChatsDir()
	if err != nil {
		return nil, fmt.Errorf("resolving chats dir: %w", err)
	}
	indexPath, err := paths.IndexPath()
	if err != nil {
		return nil, fmt.Errorf("resolving index path: %w", err)
	}
	settingsPath, err := paths.SettingsPath()
	if err != nil {
		return nil, fmt.Errorf("resolving settings path: %w", err)
	}
	cachePath, err := paths.CachePath()
	if err != nil {
		return nil, fmt.Errorf("resolving cache path: %w", err)
	}
	s.Paths = Paths{
		SessionPath:  sessionPath,
		ChatsDir:     chatsDir,
		IndexPath:    indexPath,
		SettingsPath: settingsPath,
		CachePath:    cachePath,
	}
	return &s, nil
}
