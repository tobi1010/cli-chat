package session

import (
	"cli-chat/chat"
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

func (s *Session) applySettings(set settings.Settings) {
	s.AppSettings = set
	s.Client = client.New(time.Duration(set.Timeout) * time.Second)
}

func (s *Session) applySavedState(state State) {
	if state.LastProvider.Name != "" {
		s.Provider = state.LastProvider
	}
}

func loadState(path string) (State, bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return State{}, false, nil
		}
		return State{}, false, fmt.Errorf("reading session file %s: %w", path, err)
	}
	if len(data) == 0 {
		return State{}, false, nil
	}

	var state State
	if err := json.Unmarshal(data, &state); err != nil {
		return State{}, false, fmt.Errorf("unmarshalling state: %w", err)
	}
	return State{}, true, nil

}

func (s *Session) loadChat(chatId string) error {
	if chatId == "" {
		s.Chat = chat.New()
		return nil
	}
	chatsDir, err := paths.ChatsDir()
	if err != nil {
		return fmt.Errorf("resolving chats dir: %w", err)
	}
	c, err := chat.LoadChat(chatsDir, chatId)
	if err != nil {
		s.Chat = chat.New()
		return nil
	}
	s.Chat = c
	return nil
}

func (s *Session) loadDb() error {
	indexPath, err := paths.IndexPath()
	if err != nil {
		return fmt.Errorf("resolving index path: %w", err)
	}
	db, err := index.Load(indexPath)
	if err != nil {
		return fmt.Errorf("reading index file: %w", err)
	}
	s.DB = db
	return nil
}
