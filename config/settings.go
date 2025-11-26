package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	APP_DIR  = "cli-chat"
	FILENAME = "settings.json"
)

type Settings struct {
	Provider      string `json:"provider"`
	Model         string `json:"model"`
	CommandPrefix string `json:"command_prefix"`
	Timeout       uint   `json:"timeout"`
	Columns       uint   `json:"columns"`
}

const (
	DefaultProvider      = "openai"
	DefaultModel         = "gpt-5"
	DefaultCommandPrefix = "/"
	DefaultTimeout       = 60
	DefaultColumns       = 80
)

var DefaultSettings = Settings{
	Provider:      DefaultProvider,
	Model:         DefaultModel,
	CommandPrefix: DefaultCommandPrefix,
	Timeout:       DefaultTimeout,
	Columns:       DefaultColumns,
}

func NewDefaultSettings() Settings {
	return DefaultSettings
}

func PrintSettings(path string) error {
	s, err := ReadSettings(path)
	if err != nil {
		return fmt.Errorf("error reading settings from %s: %w", path, err)
	}
	b, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(b))
	return nil
}

func ReadSettings(path string) (Settings, error) {
	file := path + "/settings.json"
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			defaults := NewDefaultSettings()
			if err := os.MkdirAll(path, 0755); err != nil {
				return Settings{}, fmt.Errorf("Error creating config directory: %w", err)
			}
			if err := WriteSettings(path, defaults); err != nil {
				return Settings{}, fmt.Errorf("Error writing settings: %w", err)
			}
			return defaults, nil
		}
		return Settings{}, fmt.Errorf("Error reading config file %w", err)
	}

	var s Settings
	err = json.Unmarshal(data, &s)
	if err != nil {
		return Settings{}, fmt.Errorf("Error unmashalling json %w", err)
	}
	return s, nil
}

func WriteSettings(path string, s Settings) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling config: %w", err)
	}
	file := path + "/settings.json"
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		return fmt.Errorf("Error writing file %w", err)
	}
	return nil
}
func getConfigPath() (string, error) {
	path, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("getting user config path: %w", err)
	}

	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working directory: %w", err)
	}
	return path, nil
}
