package core

import (
	"fmt"
	"terminal-chat/cache"
	"terminal-chat/chat"
	"terminal-chat/debug"
	"terminal-chat/index"
	"terminal-chat/internal/apitypes"
	"terminal-chat/internal/client"
	"terminal-chat/paths"
	"terminal-chat/settings"
	"time"
)

type State struct {
	LastProvider string `json:"last_provider"`
	LastModelID  string `json:"last_model_id"`
}

type Session struct {
	Registry     CommandRegistry
	Meta         []CommandMeta
	ProviderName string
	ModelID      string
	Chat         *chat.Chat
	AppSettings  settings.Settings
	DB           *index.DB
	Paths        paths.Paths
	Cache        *cache.Cache

	Client       *client.Client
	ModelLabel   string
	SystemPrompt string
}

func NewDefaultSession() (*Session, error) {
	s := Session{}

	p, err := paths.ResolvePaths()
	if err != nil {
		return nil, fmt.Errorf("resolving paths: %w", err)
	}
	s.Paths = p
	s.AppSettings = settings.NewDefaultSettings()
	s.Chat = chat.New()
	s.DB = index.NewDB()
	s.Cache = cache.New()
	s.ProviderName = s.AppSettings.DefaultProvider
	s.ModelID = apitypes.DefaultModel
	s.Client = client.New(time.Duration(60 * time.Second))
	s.ModelLabel = s.AppSettings.DefaultModel
	s.SystemPrompt = ""
	s.Registry, err = NewRegistry()
	if err != nil {
		return nil, fmt.Errorf("building registry: %w", err)
	}
	s.Meta = NewCommandMeta()

	debug.Dump("default session:", s)
	return &s, nil
}
