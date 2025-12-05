package config

import (
	"cli-chat/paths"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// EnsureSettings loads settings if present; otherwise creates defaults and writes them.
// Order: XDG_CONFIG_HOME -> $HOME/.config. No fallback to PWD.
func EnsureSettings() (Settings, string, error) {
	f, err := paths.SettingsPath()
	dir := filepath.Dir(f)

	// Try read
	data, err := os.ReadFile(f)
	if err == nil {
		var s Settings
		if uerr := json.Unmarshal(data, &s); uerr != nil {
			return Settings{}, "", fmt.Errorf("unmarshal %s: %w", f, uerr)
		}
		return s, dir, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return Settings{}, "", fmt.Errorf("read %s: %w", f, err)
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

	if err := os.WriteFile(f, encoded, 0o600); err != nil {
		return Settings{}, "", fmt.Errorf("write %s: %w", f, err)
	}

	return settings, dir, nil
}
