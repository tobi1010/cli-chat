package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Model             string `json:"model"`
	Command_delimiter string `json:"command_delimiter"`
}

func newConfig(model string) (Config, error) {
	cfg := Config{
		Model: model,
	}
	return cfg, nil
}

func readConfig(path string) (Config, error) {
	file := path + "/config"
	data, err := os.ReadFile(file)
	if err != nil {
		return Config{}, fmt.Errorf("Error reading config file %w", err)
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("Error unmashalling json %w", err)
	}
	return cfg, nil
}

func writeConfig(path string, cfg Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling config: %w", err)
	}
	file := path + "/config"
	err = os.WriteFile(file, data, 0664)
	if err != nil {
		return fmt.Errorf("Error writing file %w", err)
	}
	return nil
}
