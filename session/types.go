package session

import (
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
	LastProvider providers.Provider `json:"lastProvider"`
}
type Paths struct {
	SessionPath  string `json:"session_path"`
	ChatsDir     string `json:"chats_dir"`
	IndexPath    string `json:"index_path"`
	SettingsPath string `json:"settings_path"`
}

type Session struct {
	Provider    providers.Provider
	Chat        *chat.Chat
	Client      *client.Client
	AppSettings settings.Settings
	DB          *index.DB
	Paths       Paths
}

func NewDefaultSession() (*Session, error) {
	s := Session{}
	s.Provider = providers.NewDefault()
	s.AppSettings = settings.NewDefaultSettings()
	s.Client = client.New(time.Duration(s.AppSettings.Timeout) * time.Second)
	s.Chat = chat.New()
	s.DB = index.NewDB()

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
	s.Paths = Paths{
		SessionPath:  sessionPath,
		ChatsDir:     chatsDir,
		IndexPath:    indexPath,
		SettingsPath: settingsPath,
	}
	return &s, nil
}
