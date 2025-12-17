package commands

import (
	"cli-chat/paths"
	"cli-chat/session"
	"fmt"
)

func CommandSetPrefix(s *session.Session, args []string) error {
	s.AppSettings.CommandPrefix = args[0]
	settingsPath, err := paths.SessionPath()
	if err != nil {
		return fmt.Errorf("resolving settings path: %w", err)
	}
	if err = s.AppSettings.Save(settingsPath); err != nil {
		return fmt.Errorf("writingSettings: %w", err)
	}
	return nil
}
