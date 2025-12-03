package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetConfigRoot() (string, error) {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return xdg, nil
	}
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return "", fmt.Errorf("cannot resolve config root: set XDG_CONFIG_HOME or ensure a home directory is available: %w", err)
	}
	return filepath.Join(home, ".config"), nil
}

func appDir(root string) string { return filepath.Join(root, APP_DIR) }

func settingsPath(root string) string {
	return filepath.Join(root, APP_DIR, FILENAME)
}
