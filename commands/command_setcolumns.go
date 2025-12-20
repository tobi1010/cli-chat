package commands

import (
	"cli-chat/session"
	"fmt"
	"strconv"
)

func CommandSetColumns(s *session.Session, args []string) error {
	length, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("argument must be a number!")
	}
	if length < 5 {
		fmt.Println("please set cloums to a value > 5")
		return nil
	}
	s.AppSettings.Columns = length
	if err := s.AppSettings.Save(s.Paths.SettingsPath); err != nil {
		return fmt.Errorf("writing settings: %w", err)
	}
	return nil
}
