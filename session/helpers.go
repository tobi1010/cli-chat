package session

import (
	"cli-chat/chat"
	"cli-chat/index"
	"cli-chat/internal/client"
	"cli-chat/providers"
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

func (s *Session) applySavedState(state State) error {
	if state.LastProvider != "" {
		lastProv, err := providers.New(*s.Cache, state.LastProvider, state.LastModelID)
		if err != nil {
			return fmt.Errorf(
				"resolving last provider %q with model %q: %w",
				state.LastProvider,
				state.LastModelID,
				err,
			)
		}
		s.Provider = lastProv
	}
	return nil
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
	return state, true, nil

}

func (s *Session) loadChat(chatId string) error {
	if chatId == "" {
		s.Chat = chat.New()
		return nil
	}
	c, err := chat.Load(s.Paths.ChatsDir, chatId)
	if err != nil {
		return fmt.Errorf("load chat %q: %w", chatId, err)
	}
	s.Chat = c
	return nil
}

func (s *Session) loadDb() error {
	db, err := index.Load(s.Paths.IndexPath)
	if err != nil {
		return fmt.Errorf("reading index file: %w", err)
	}
	s.DB = db
	return nil
}

func (s *Session) toState() State {
	st := State{
		LastProvider: s.Provider.Name,
		LastModelID:  s.Provider.Model.ID,
	}
	return st

}
