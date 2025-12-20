package commands

import (
	"cli-chat/chat"
	"cli-chat/session"
	"fmt"
)

func CommandNew(s *session.Session, args []string) error {
	s.Chat = chat.New()
	if err := s.Save(s.Paths.SessionPath); err != nil {
		return fmt.Errorf("saving session: %w", err)
	}
	return nil
}
