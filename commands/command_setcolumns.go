package commands

import (
	"cli-chat/paths"
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
	settingsPath, err := paths.SettingsPath()
	if err != nil {
		return fmt.Errorf("resolving session path: %w", err)
	}
	err = s.AppSettings.Save(settingsPath)
	if err != nil {
		return fmt.Errorf("writing settings: %w", err)
	}
	return nil
}
