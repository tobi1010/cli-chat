package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// EnsureSettings loads settings if present; otherwise creates defaults and writes them.
// Order: XDG_CONFIG_HOME -> $HOME/.config. No fallback to PWD.
func EnsureSettings() (Settings, string, error) {
	root, err := GetConfigRoot()
	if err != nil {
		return Settings{}, "", err
	}
	dir := appDir(root)
	path := settingsPath(root)

	// Try read
	data, err := os.ReadFile(path)
	if err == nil {
		var s Settings
		if uerr := json.Unmarshal(data, &s); uerr != nil {
			return Settings{}, "", fmt.Errorf("unmarshal %s: %w", path, uerr)
		}
		return s, dir, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return Settings{}, "", fmt.Errorf("read %s: %w", path, err)
	}

	// Not found: create dir and write defaults
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return Settings{}, "", fmt.Errorf("mkdir %s: %w", dir, err)
	}

	settings := NewDefaultSettings()
	encoded, merr := json.MarshalIndent(settings, "", "  ")
	if merr != nil {
		return Settings{}, "", fmt.Errorf("marshal defaults: %w", merr)
	}

	if err := os.WriteFile(path, encoded, 0o600); err != nil {
		return Settings{}, "", fmt.Errorf("write %s: %w", path, err)
	}

	return settings, dir, nil
}
