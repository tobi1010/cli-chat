package commands

import (
	"fmt"
	"strconv"
	"terminal-chat/session"
)

func CommandSetTTL(s *session.Session, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: /set-ttl <seconds>")
	}
	ttl, err := strconv.Atoi(args[0])
	if err != nil || ttl < 0 {
		return fmt.Errorf("usage: /set-ttl <seconds>")
	}
	s.AppSettings.TTL = ttl
	if err := s.AppSettings.Save(s.Paths.SettingsPath); err != nil {
		return fmt.Errorf("writingSettings: %w", err)
	}
	return nil
}
