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
	LastChatId   string             `json:"lastChatId"`
	LastProvider providers.Provider `json:"lastProvider"`
}

type Session struct {
	Provider    providers.Provider
	Chat        *chat.Chat
	Client      *client.Client
	AppSettings settings.Settings
	DB          *index.DB
}

func NewDefaultSession() (*Session, error) {
	s := Session{}
	s.Provider = providers.NewDefault()
	s.AppSettings = settings.NewDefaultSettings()
	s.Client = client.New(time.Duration(s.AppSettings.Timeout) * time.Second)
	s.Chat = chat.New()

	indexPath, err := paths.IndexPath()
	if err != nil {
		return nil, fmt.Errorf("resolving index path: %w", err)
	}
	chatsDir, err := paths.ChatsDir()
	if err != nil {
		return nil, fmt.Errorf("resolving chats dir: %w", err)
	}
	s.DB, err = index.NewDB(indexPath, chatsDir)
	if err != nil {
		return nil, fmt.Errorf("loading db: %w", err)
	}
	return &s, nil
}
