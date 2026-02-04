package core

import (
	"fmt"
	"terminal-chat/chat"
)

func CommandNew(s *Session, args []string) error {
	s.Chat = chat.New()
	if err := s.SessionSave(s.Paths.SessionPath); err != nil {
		return fmt.Errorf("saving session: %w", err)
	}
	return nil
}
