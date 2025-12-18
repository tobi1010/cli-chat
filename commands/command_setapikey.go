package commands

import (
	"cli-chat/paths"
	"cli-chat/session"
	"fmt"
)

func CommandSetApiKey(s *session.Session, args []string) error {
	if args[0] != "" {
		s.AppSettings.ApiKey = args[0]
	}
	path, err := paths.SessionPath()
	if err != nil {
		return fmt.Errorf("resolving session path: %w", err)
	}
	if err := s.Save(path); err != nil {
		return fmt.Errorf("saving session to %s: %w", path, err)
	}
	return nil
}
