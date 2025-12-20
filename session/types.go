package session

import (
	"cli-chat/chat"
	"cli-chat/index"
	"cli-chat/internal/client"
	"cli-chat/providers"
	"cli-chat/settings"
	"time"
)

type State struct {
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
	s.DB = index.NewDB()
	return &s, nil
}
