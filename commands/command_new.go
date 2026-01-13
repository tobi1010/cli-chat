package commands

import (
	"fmt"
	"terminal-chat/chat"
	"terminal-chat/session"
)

func CommandNew(s *session.Session, args []string) error {
	s.Chat = chat.New()
	if err := s.Save(s.Paths.SessionPath); err != nil {
		return fmt.Errorf("saving session: %w", err)
	}
	return nil
}
