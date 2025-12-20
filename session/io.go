package session

import (
	"cli-chat/fileatomic"
	"cli-chat/settings"
	"encoding/json"
	"fmt"
)

func LoadOrCreate(sessionPath string) (*Session, error) {
	s, err := NewDefaultSession()
	if err != nil {
		return nil, fmt.Errorf("creating default session: %w", err)
	}

	set, err := settings.LoadOrCreate()
	if err != nil {
		return nil, fmt.Errorf("load or create settings: %w", err)
	}
	s.applySettings(set)

	state, ok, err := loadState(sessionPath)
	if err != nil {
		return nil, fmt.Errorf("loading session file: %w", err)
	}
	if ok {
		s.applySavedState(state)
	}

	if err := s.loadDb(); err != nil {
		return nil, fmt.Errorf("loading database: %w", err)
	}

	lastChatId := s.DB.GetLastChatId()
	if err := s.loadChat(lastChatId); err != nil {
		return nil, fmt.Errorf("loading chat: %w", err)
	}

	return s, nil
}

func (s *Session) Save(sessionPath string) error {
	st := State{
		LastProvider: s.Provider,
	}
	data, err := json.Marshal(st)
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	if err = fileatomic.Write(sessionPath, data, 0o600); err != nil {
		return fmt.Errorf("writing file atomically %s: %w", sessionPath, err)
	}

	return nil
}
