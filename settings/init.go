package settings

import (
	"cli-chat/fileatomic"
	"cli-chat/paths"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// EnsureSettings loads settings if present; otherwise creates defaults and writes them.
// Order: XDG_CONFIG_HOME -> $HOME/.config. No fallback to PWD.
func EnsureSettings() (Settings, error) {
	f, err := paths.SettingsPath()
	dir := filepath.Dir(f)

	// Try read
	data, err := os.ReadFile(f)
	if err == nil {
		var s Settings
		if err := json.Unmarshal(data, &s); err != nil {
			return Settings{}, fmt.Errorf("unmarshal %s: %w", f, err)
		}
		return s, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return Settings{}, fmt.Errorf("read %s: %w", f, err)
	}

	// Not found: create dir and write defaults
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return Settings{}, fmt.Errorf("mkdir %s: %w", dir, err)
	}

	defSettings := NewDefaultSettings()
	encoded, err := json.MarshalIndent(defSettings, "", "  ")
	if err != nil {
		return Settings{}, fmt.Errorf("marshal defaults: %w", err)
	}

	if err = fileatomic.Write(f, encoded, 0o600); err != nil {
		return Settings{}, fmt.Errorf("write %s: %w", f, err)
	}

	return defSettings, nil
}
