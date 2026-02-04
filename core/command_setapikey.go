package core

import (
	"fmt"
)

func CommandSetApiKey(s *Session, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: %sset-apikey <apikey name>", s.AppSettings.CommandPrefix)
	}
	if args[0] != "" {
		s.AppSettings.ApiKey = args[0]
	}
	if err := s.SessionSave(s.Paths.SessionPath); err != nil {
		return fmt.Errorf("saving session to %s: %w", s.Paths.SessionPath, err)
	}
	return nil
}
