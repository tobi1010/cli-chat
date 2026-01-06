package session

import (
	"cli-chat/cache"
	"cli-chat/chat"
	"cli-chat/index"
	"cli-chat/internal/apitypes"
	"cli-chat/internal/client"
	"cli-chat/paths"
	"cli-chat/providers"
	"cli-chat/settings"
	"time"
)

type State struct {
	LastProvider string `json:"last_provider"`
	LastModelID  string `json:"last_model_id"`
}

type Session struct {
	ProviderName string
	ModelID      string
	Chat         *chat.Chat
	AppSettings  settings.Settings
	DB           *index.DB
	Paths        paths.Paths
	Cache        *cache.Cache

	Client *client.Client
}

func NewDefaultSession() (*Session, error) {
	s := Session{}

	s.AppSettings = settings.NewDefaultSettings()
	s.Chat = chat.New()
	s.DB = index.NewDB()
	s.Cache = cache.New()
	s.ProviderName = providers.Default
	s.ModelID = apitypes.DefaultModel
	s.Client = client.New(time.Duration(60 * time.Second))

	return &s, nil
}
