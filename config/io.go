package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func PrintSettings() error {
	path, err := GetSettingsPath()
	if err != nil {
		return fmt.Errorf("getting config path: %w", err)
	}
	var s Settings
	err = ReadSettings(&s)
	if err != nil {
		return fmt.Errorf("reading settings from %s: %w", path, err)
	}
	b, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(b))
	return nil
}

func ReadSettings(settings *Settings) error {
	// InitDefaultSettings must be called first to ensure dir exists

	path, err := GetSettingsPath()
	if err != nil {
		return fmt.Errorf("getting config path: %w", err)
	}
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

func InitDefaultSettings() error {
	// must be called on startup to guarantee APP_DIR exitsts
	path, err := GetSettingsPath()
	if err != nil {
		return fmt.Errorf("resolving config path: %w", err)
	}
	fp := filepath.Join(path, FILENAME)
	if info, err := os.Stat(fp); err == nil && info.Mode().IsRegular() {
		return nil
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("stat %s: %w", fp, err)
	}
	if err := os.MkdirAll(path, 0o700); err != nil {
		return fmt.Errorf("creating config dir: %w", err)
	}
	defaults := NewDefaultSettings()
	data, err := json.MarshalIndent(defaults, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	if err := os.WriteFile(fp, data, 0o600); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}
	return nil
}

func WriteSettings(s *Settings) error {
	// InitDefaultSettings must be called first to ensunre dir exists
	path, err := GetSettingsPath()
	if err != nil {
		return fmt.Errorf("getting settings path: %w", err)
	}
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
