package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func PrintSettings(path string) error {
	var s Settings
	err := ReadSettings(&s, path)
	if err != nil {
		return fmt.Errorf("reading settings from %s: %w", path, err)
	}
	b, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(b))
	return nil
}

func ReadSettings(settings *Settings, path string) error {
	file := filepath.Join(path, FILENAME)
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("reading config file: %w", err)
	}

	err = json.Unmarshal(data, settings)
	if err != nil {
		return fmt.Errorf("unmarshalling json: %w", err)
	}
	return nil
}

func WriteSettings(s *Settings, path string) error {
	data, err := json.MarshalIndent(*s, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling config: %w", err)
	}
	file := filepath.Join(path, FILENAME)
	if err = os.WriteFile(file, data, 0o600); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}
	return nil
}
