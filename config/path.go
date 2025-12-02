package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetSettingsPath() (string, error) {
	xdg := os.Getenv("XDG_CONFIG_HOME")
	if xdg != "" {
		return filepath.Join(xdg, APP_DIR), nil
	}
	// ucd, err := os.UserConfigDir()
	// if err != nil {
	// 	return "", fmt.Errorf("resolving user config dir: %w", err)
	// }
	// if ucd != "" {
	// 	return filepath.Join(ucd, APP_DIR), nil
	// }
	home := os.Getenv("HOME")
	if home != "" {
		return filepath.Join(home, ".config", APP_DIR), nil
	}
	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("resolving pwd: %w", err)
	}
	log.Println("could not resolve config dir, falling back to pwd")
	return pwd, nil
}
