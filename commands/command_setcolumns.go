package commands

import (
	"fmt"
	"strconv"
	"terminal-chat/session"
)

func CommandSetColumns(s *session.Session, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: %set-columns <columns>", s.AppSettings.CommandPrefix)
	}
	length, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("argument must be a number!")
	}
	if length < 5 {
		return fmt.Errorf("please set cloums to a value > 5")
	}
	s.AppSettings.Columns = length
	if err := s.AppSettings.Save(s.Paths.SettingsPath); err != nil {
		return fmt.Errorf("writing settings: %w", err)
	}
	return nil
}
