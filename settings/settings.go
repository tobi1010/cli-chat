package settings

import (
	"encoding/json"
	"fmt"
	"os"
)

type Settings struct {
	Provider      string `json:"provider"`
	Model         string `json:"model"`
	CommandPrefix string `json:"command_prefix"`
	Timeout       int    `json:"timeout"`
}

func Print(path string) error {
	s, err := Read(path)
	if err != nil {
		return fmt.Errorf("error reading settings from %s: %w", path, err)
	}
	b, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(b))
	return nil
}

func Read(path string) (Settings, error) {
	file := path + "/settings.json"
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			defaults := Settings{
				Provider:      "openai",
				Model:         "gpt-5",
				CommandPrefix: "/",
				Timeout:       60,
			}
			if err := os.MkdirAll(path, 0755); err != nil {
				return Settings{}, fmt.Errorf("Error creating config directory: %w", err)
			}
			if err := Write(path, defaults); err != nil {
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

func Write(path string, s Settings) error {
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
