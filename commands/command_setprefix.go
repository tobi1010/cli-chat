package commands

import (
	"cli-chat/config"
	"fmt"
)

func CommandSetPrefix(cfg *config.Config, args []string) error {
	cfg.AppSettings.CommandPrefix = args[0]
	err := cfg.AppSettings.Save()
	if err != nil {
		return fmt.Errorf("writingSettings: %w", err)
	}
	return nil
}
