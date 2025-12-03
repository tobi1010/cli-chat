package commands

import (
	"cli-chat/config"
	"fmt"
)

func CommandSetPrefix(cfg *config.Config, args []string) error {
	cfg.AppSettings.CommandPrefix = args[0]
	err := config.WriteSettings(cfg.AppSettings, cfg.SettingsPath)
	if err != nil {
		return fmt.Errorf("writingSettings: %w", err)
	}
	return nil
}
