package main

import (
	"cli-chat/config"
	"fmt"
)

func commandPrintSettings(cfg *config.Config, args []string) error {
	err := config.PrintSettings()
	if err != nil {
		return fmt.Errorf("printig settings: %w", err)
	}
	return nil
}
