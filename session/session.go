package session

import (
	"cli-chat/chat"
	"cli-chat/fileatomic"
	"cli-chat/index"
	"cli-chat/internal/client"
	"cli-chat/paths"
	"cli-chat/settings"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type State struct {
	LastChatId   string   `json:"lastChatId"`
	LastProvider Provider `json:"lastProvider"`
	LastModel    string   `json:"lastModel"`
}

type Provider struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	BaseURL string `json:"baseurl"`
}

func NewDefaultSession() *Session {
	s := Session{
		Provider: Provider{
			Name:    "openai",
			Key:     "OPENAI_API_KEY",
			BaseURL: "https://api.openai.com/v1/",
		},
		Model: "gpt-5",
	}
	settings := settings.NewDefaultSettings()
	s.AppSettings = settings
	return &s
}

type Session struct {
	Provider    Provider
	Model       string
	Chat        *chat.Chat
	Client      *client.Client
	AppSettings settings.Settings
	DB          *index.DB
}

func LoadOrCreate() (*Session, error) {
	sessionPath, err := paths.SessionPath()
	if err != nil {
		return nil, fmt.Errorf("resolving session path: %w", err)
	}
	indexPath, err := paths.IndexPath()
	if err != nil {
		return nil, fmt.Errorf("resolving index path: %w", err)
	}
	s := NewDefaultSession()
	s.Client = client.New(60 * time.Second)
	set, err := settings.EnsureSettings()
	if err != nil {
		return nil, fmt.Errorf("ensure settings: %w", err)
	}
	s.AppSettings = set

	data, err := os.ReadFile(sessionPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("reading session file %s: %w", sessionPath, err)
	}
	var state State
	_ = json.Unmarshal(data, &state)
	if state.LastChatId == "" {
		s.Chat = chat.New()
	} else {
		s.Chat, err = chat.ReadChat(state.LastChatId)
		if err != nil {
			s.Chat = chat.New()
		}
	}
	db, err := index.Load(indexPath)
	if err != nil {
		return nil, fmt.Errorf("reading index file: %w", err)
	}
	s.DB = db
	if state.LastProvider.Name == "" {
		return s, nil
	}
	s.Provider = state.LastProvider
	if state.LastModel != "" {
		s.Model = state.LastModel
	}
	return s, nil
}

func (s *Session) Save() error {
	st := State{
		LastChatId:   s.Chat.ID,
		LastProvider: s.Provider,
		LastModel:    s.Model,
	}
	data, err := json.Marshal(st)
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	sessionPath, err := paths.SessionPath()
	if err != nil {
		return fmt.Errorf("resolving session path: %w", err)
	}
	err = fileatomic.Write(sessionPath, data, 0o600)
	if err != nil {
		return fmt.Errorf("writing file atomically %s: %w", sessionPath, err)
	}

	return nil
}
