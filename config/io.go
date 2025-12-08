package config

import (
	"cli-chat/fileatomic"
	"cli-chat/paths"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func PrintSettings() error {
	s, err := ReadSettings()
	if err != nil {
		return fmt.Errorf("reading settings: %w", err)
	}
	b, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(b))
	return nil
}

func ReadSettings() (*Settings, error) {
	f, err := paths.SettingsPath()
	if err != nil {
		return nil, fmt.Errorf("resolving settings path: %w", err)
	}
	data, err := os.ReadFile(f)
	var settings Settings
	if err != nil {
		if os.IsNotExist(err) {
			settings = NewDefaultSettings()
			return &settings, nil
		}
		return nil, fmt.Errorf("reading config file: %w", err)
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling json: %w", err)
	}
	return &settings, nil
}

func (s *Settings) Save() error {
	path, err := paths.SettingsPath()
	if err != nil {
		return fmt.Errorf("resolving settings path: %w", err)
	}
	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, 0o700)
	if err != nil {
		return fmt.Errorf("creating settings dir: %w", err)
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling config: %w", err)
	}
	err = fileatomic.Write(path, data, 0o600)
	if err != nil {
		return fmt.Errorf("writing settings: %w", err)
	}
	return nil
}
