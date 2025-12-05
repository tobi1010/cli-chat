package config

import (
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
	f, err := paths.SettingsPath()
	if err != nil {
		return fmt.Errorf("resolving settings path: %w", err)
	}
	dir := filepath.Dir(f)
	err = os.MkdirAll(dir, 0o700)
	if err != nil {
		return fmt.Errorf("creating settings dir: %w", err)
	}
	tmpFile := f + ".tmp"

	fd, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return fmt.Errorf("opening temp file: %w", err)
	}

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		_ = fd.Close()
		return fmt.Errorf("marshalling config: %w", err)
	}

	_, err = fd.Write(data)
	if err != nil {
		_ = fd.Close()
		return fmt.Errorf("writing data: %w", err)
	}
	err = fd.Sync()
	if err != nil {
		_ = fd.Close()
		return fmt.Errorf("syncing file: %w", err)
	}
	err = fd.Close()
	if err != nil {
		return fmt.Errorf("closing file: %w", err)
	}
	err = os.Rename(tmpFile, f)
	if err != nil {
		_ = os.Remove(tmpFile)
		return fmt.Errorf("renaming file: %w", err)
	}
	return nil
}
