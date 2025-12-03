package commands

import (
	"cli-chat/config"
	"fmt"
)

func CommandPrintSettings(cfg *config.Config, args []string) error {
	err := config.PrintSettings(cfg.SettingsPath)
	if err != nil {
		return fmt.Errorf("printig settings: %w", err)
	}
	return nil
}
