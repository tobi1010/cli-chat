package session

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"terminal-chat/chat"
	"terminal-chat/debug"
	"terminal-chat/index"
	"terminal-chat/paths"
	"terminal-chat/settings"
)

func applyDb(s *Session, paths paths.Paths) error {
	db, err := index.Load(paths.IndexPath)
	if err != nil {
		//create new db
		db = index.NewDB()
		s.DB = db
		if err := db.Save(paths.IndexPath); err != nil {
			return fmt.Errorf("saving db file: %w", err)
		}
	}
	s.DB = db
	return nil
}

func applySettings(s *Session, paths paths.Paths) error {
	set, err := settings.Load(paths.SettingsPath)
	if err != nil {
		//apply default settings
		set = settings.NewDefaultSettings()
		s.AppSettings = set
		err := s.AppSettings.Save(paths.SettingsPath)
		if err != nil {
			return fmt.Errorf("saving settings file: %w", err)
		}
	}
	debug.Dump(set)
	s.AppSettings = set
	s.AppSettings.Save(paths.SettingsPath)
	return nil
}

func resolveChat(db index.DB, paths paths.Paths) (*chat.Chat, error) {
	lastChatID := db.GetLastChatId()
	if lastChatID == "" {
		return chat.New(), nil
	} else {
		c, err := chat.Load(paths.ChatsDir, lastChatID)
		if err != nil {
			return nil, fmt.Errorf("reading last chat, id: %s: %w", lastChatID, err)
		}
		return c, nil
	}
}

func loadState(path string) (State, bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return State{}, false, nil
		}
		return State{}, false, fmt.Errorf("reading state file %s: %w", path, err)
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

func (s *Session) loadDb() error {
	db, err := index.Load(s.Paths.IndexPath)
	if err != nil {
		return fmt.Errorf("reading index file: %w", err)
	}
	s.DB = db
	return nil
}
