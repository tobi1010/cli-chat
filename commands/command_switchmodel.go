package commands

import (
	"cli-chat/config"
	"fmt"
)

func CommandSwitchModel(cfg *config.Config, args []string) error {
	cfg.AppSettings.Model = args[0]
	err := config.WriteSettings(cfg.AppSettings)
	if err != nil {
		return fmt.Errorf("writing settings.json: %w", err)
	}
	return nil
}
