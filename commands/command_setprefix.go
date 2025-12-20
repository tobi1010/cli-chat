package commands

import (
	"cli-chat/session"
	"fmt"
)

func CommandSetPrefix(s *session.Session, args []string) error {
	s.AppSettings.CommandPrefix = args[0]
	if err := s.AppSettings.Save(s.Paths.SettingsPath); err != nil {
		return fmt.Errorf("writingSettings: %w", err)
	}
	return nil
}
