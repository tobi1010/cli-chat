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
	Provider      Provider `json:"provider"`
	Model         string   `json:"model"`
	CommandPrefix string   `json:"command_prefix"`
	Timeout       int      `json:"timeout"`
	Columns       int      `json:"columns"`
}
type Provider struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	BaseURL string `json:"baseurl"`
}

var DefaultProvider = Provider{
	Name:    "openai",
	Key:     "OPENAI_API_KEY",
	BaseURL: "https://api.openai.com/v1/",
}

const (
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
		return fmt.Errorf("reading settings from %s: %w", path, err)
	}
	b, _ := json.MarshalIndent(s, "", "  ")
	fmt.Println(string(b))
	return nil
}

func ReadSettings(settings *Settings) error {
	// InitDefaultSettings must be called first to ensunre dir exists

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
	fmt.Printf("settings path: %s", path)
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
