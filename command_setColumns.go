package main

import (
	"cli-chat/config"
	"fmt"
)

func commandSetColumns(cfg *config.Config, length uint) error {
	cfg.AppSettings.Columns = length
	err := config.WriteSettings(cfg.AppSettings)
	return nil
}
