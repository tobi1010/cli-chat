package commands

import (
	"fmt"
	"terminal-chat/session"
)

func CommandPrintSettings(s *session.Session, args []string) error {
	set := s.AppSettings
	if err := set.PrintSettings(); err != nil {
		return fmt.Errorf("printig settings: %w", err)
	}
	return nil
}
