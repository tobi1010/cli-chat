package commands

import (
	"cli-chat/session"
	"fmt"
)

func CommandPrintSettings(s *session.Session, args []string) error {
	set := s.AppSettings
	if err := set.PrintSettings(); err != nil {
		return fmt.Errorf("printig settings: %w", err)
	}
	return nil
}
