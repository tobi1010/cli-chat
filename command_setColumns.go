package main

import (
	"cli-chat/config"
	"fmt"
	"strconv"
)

func commandSetColumns(cfg *config.Config, args []string) error {
	length, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("argument must be a number!")
	}
	if length < 5 {
		fmt.Println("please set cloums to a value > 5")
		return nil
	}
	cfg.AppSettings.Columns = length
	err = config.WriteSettings(cfg.AppSettings)
	if err != nil {
		return fmt.Errorf("writing settings: %w", err)
	}
	return nil
}
