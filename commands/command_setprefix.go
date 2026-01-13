package commands

import (
	"fmt"
	"terminal-chat/session"
)

func CommandSetPrefix(s *session.Session, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: %sset-prefix <prefix>", s.AppSettings.CommandPrefix)
	}
	if len(args[0]) > 3 {
		return fmt.Errorf("usage: %sset-prefix <prefix> (max 3 chars)", s.AppSettings.CommandPrefix)
	}
	s.AppSettings.CommandPrefix = args[0]
	if err := s.AppSettings.Save(s.Paths.SettingsPath); err != nil {
		return fmt.Errorf("writingSettings: %w", err)
	}
	return nil
}
