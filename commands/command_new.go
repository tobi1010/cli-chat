package commands

import (
	"cli-chat/chat"
	"cli-chat/paths"
	"cli-chat/session"
	"fmt"
)

func CommandNew(s *session.Session, args []string) error {
	sessionPath, err := paths.SessionPath()
	if err != nil {
		return fmt.Errorf("resolving session path: %w", err)
	}
	s.Chat = chat.New()
	if err = s.Save(sessionPath); err != nil {
		return fmt.Errorf("saving session: %w", err)
	}
	return nil
}
