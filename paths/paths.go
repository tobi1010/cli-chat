package paths

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	AppDirName   = "cli-chat"
	SettingsFile = "settings.json"
	IndexFile    = "index.json"
	ChatsDirName = "chats"
	SessionFile  = "session.json"
	CacheFile    = "modelcache.json"
)

// returns the base config root (XDG_CONFIG_HOME or ~/.config).
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

// returns the app-specific config directory.
func AppConfigDir() (string, error) {
	root, err := GetConfigRoot()
	if err != nil {
		return "", err
	}
	return filepath.Join(root, AppDirName), nil
}

// returns the full path to settings.json.
func SettingsPath() (string, error) {
	dir, err := AppConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, SettingsFile), nil
}

// returns the app-specific data directory (XDG_DATA_HOME or ~/.local/share/AppDirName).
func AppDataDir() (string, error) {
	if xdg := os.Getenv("XDG_DATA_HOME"); xdg != "" {
		return filepath.Join(xdg, AppDirName), nil
	}
	home, err := os.UserHomeDir()
	if err != nil || home == "" {
		return "", fmt.Errorf("resolving user home dir: %w", err)
	}
	return filepath.Join(home, ".local", "share", AppDirName), nil
}

func SessionPath() (string, error) {
	dataDir, err := AppDataDir()
	if err != nil {
		return "", fmt.Errorf("resolving app data dir: %w", err)
	}
	return filepath.Join(dataDir, SessionFile), nil

}

// returns the full path to the index file (db).
func IndexPath() (string, error) {
	dataDir, err := AppDataDir()
	if err != nil {
		return "", fmt.Errorf("resolving app data dir: %w", err)
	}
	return filepath.Join(dataDir, IndexFile), nil
}

// returns the directory for chat files.
func ChatsDir() (string, error) {
	dataDir, err := AppDataDir()
	if err != nil {
		return "", fmt.Errorf("resolving app data dir: %w", err)
	}
	return filepath.Join(dataDir, ChatsDirName), nil
}

// returns the full path of the cache file
func CachePath() (string, error) {
	dataDir, err := AppDataDir()
	if err != nil {
		return "", fmt.Errorf("resolving app data dir: %w", err)
	}
	return filepath.Join(dataDir, CacheFile), nil
}

func StatePath() (string, error) {
	dataDir, err := AppDataDir()
	if err != nil {
		return "", fmt.Errorf("resolving app data dir: %w", err)
	}
	return dataDir, nil
}
