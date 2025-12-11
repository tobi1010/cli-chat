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

type PersistedState struct {
	LastChatId   string   `json:"lastChatId"`
	LastProvider Provider `json:"lastProvider"`
	LastModel    string   `json:"lastModel"`
}

type Provider struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	BaseURL string `json:"baseurl"`
}

type Session struct {
	Provider    Provider
	Model       string
	Chat        *chat.Chat
	Client      *client.Client
	AppSettings settings.Settings
	DB          *index.DB
}

func NewDefaultSession() (*Session, error) {
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
	s.Client = client.New(time.Duration(settings.Timeout) * time.Second)
	s.Chat = chat.New()
	indexPath, err := paths.IndexPath()
	if err != nil {
		return nil, fmt.Errorf("resolving index path: %w", err)
	}
	s.DB, err = index.Load(indexPath)
	if err != nil {
		return nil, fmt.Errorf("loading db: %w", err)
	}
	return &s, nil
}

func LoadOrCreate(sessionPath string) (*Session, error) {
	indexPath, err := paths.IndexPath()
	if err != nil {
		return nil, fmt.Errorf("resolving index path: %w", err)
	}
	s, err := NewDefaultSession()
	if err != nil {
		return nil, fmt.Errorf("creating default session: %w", err)
	}

	// load stored values
	set, err := settings.EnsureSettings()
	if err != nil {
		return nil, fmt.Errorf("ensure settings: %w", err)
	}
	s.AppSettings = set

	data, err := os.ReadFile(sessionPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("reading session file %s: %w", sessionPath, err)
	}
	var state PersistedState
	_ = json.Unmarshal(data, &state)
	if state.LastChatId == "" {
		s.Chat = chat.New()
	} else {
		chatsDir, err := paths.ChatsDir()
		if err != nil {
			return nil, fmt.Errorf("resolving chats dir: %w", err)
		}
		s.Chat, err = chat.ReadChat(chatsDir, state.LastChatId)
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

func (s *Session) Save(sessionPath string) error {
	st := PersistedState{
		LastChatId:   s.Chat.ID,
		LastProvider: s.Provider,
		LastModel:    s.Model,
	}
	data, err := json.Marshal(st)
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	err = fileatomic.Write(sessionPath, data, 0o600)
	if err != nil {
		return fmt.Errorf("writing file atomically %s: %w", sessionPath, err)
	}

	return nil
}
