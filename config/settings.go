package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	APP_DIR  = "cli-chat"
	FILENAME = "settings.json"
)

type Settings struct {
	Provider      string `json:"provider"`
	Model         string `json:"model"`
	CommandPrefix string `json:"command_prefix"`
	Timeout       int    `json:"timeout"`
	Columns       int    `json:"columns"`
}

const (
	DefaultProvider      = "openai"
	DefaultModel         = "gpt-5"
	DefaultCommandPrefix = "/"
	DefaultTimeout       = 60
	DefaultColumns       = 80
)

var DefaultSettings = Settings{
	Provider:      DefaultProvider,
	Model:         DefaultModel,
	CommandPrefix: DefaultCommandPrefix,
	Timeout:       DefaultTimeout,
	Columns:       DefaultColumns,
}

func NewDefaultSettings() Settings {
	return DefaultSettings
}

func PrintSettings() error {
	path, err := GetSettingsPath()
	if err != nil {
		return fmt.Errorf("getting config path: %w", err)
	}
	var s Settings
	err = ReadSettings(&s)
	if err != nil {
		return fmt.Errorf("error reading settings from %s: %w", path, err)
	}
	b, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(b))
	return nil
}

func ReadSettings(settings *Settings) error {
	path, err := GetSettingsPath()
	if err != nil {
		return fmt.Errorf("getting config path: %w", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Error reading config file %w", err)
	}

	err = json.Unmarshal(data, settings)
	if err != nil {
		return fmt.Errorf("Error unmashalling json %w", err)
	}
	return nil
}

func getConfigPath() (string, error) {
	if xdg := os.Getenv("XDG_CONFIG_HOME"); xdg != "" {
		return filepath.Join(xdg, APP_DIR), nil
	}
	if xdg := os.Getenv("XDG_HOME"); xdg != "" {
		return filepath.Join(xdg, ".config", APP_DIR), nil
	}
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolving user home dir: %w", err)
	}
	return filepath.Join(userHome, ".config", APP_DIR), nil
}

func InitDefaultSettings() error {
	path, err := getConfigPath()
	if err != nil {
		return fmt.Errorf("resolving config path: %w", err)
	}
	fp := filepath.Join(path, FILENAME)
	if info, err := os.Stat(fp); err == nil && info.Mode().IsRegular() {
		return nil
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("stat: %s, %w", info, err)
	}
	if err := os.MkdirAll(path, 0700); err != nil {
		return fmt.Errorf("creating config dir: %w", err)
	}
	defaults := NewDefaultSettings()
	data, err := json.MarshalIndent(defaults, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	if err := os.WriteFile(fp, data, 0600); err != nil {
		return fmt.Errorf("writing file: %w", err)
	}
	return nil
}

func WriteSettings(s *Settings) error {
	path, err := GetSettingsPath()
	if err != nil {
		return fmt.Errorf("getting settings path: %w", err)
	}
	data, err := json.MarshalIndent(*s, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshalling config: %w", err)
	}
	file := path
	err = os.WriteFile(file, data, 0600)
	if err != nil {
		return fmt.Errorf("Error writing file %w", err)
	}
	return nil
}
func GetSettingsPath() (string, error) {
	configPath, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("getting user config path: %w", err)
	}
	path := filepath.Join(configPath, APP_DIR, FILENAME)
	found, err := fileExists(path)
	if err != nil {
		return "", fmt.Errorf("checking settings file: %w", err)
	}
	if found {
		return path, nil
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("getting working directory: %w", err)
	}
	path = filepath.Join(pwd, FILENAME)
	return path, nil
}

func fileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, fmt.Errorf("stat %s: %w", path, err)
	}
	return info.Mode().IsRegular(), nil
}
