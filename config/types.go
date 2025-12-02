package config

import (
	"cli-chat/internal/client"
	"time"
)

const (
	DefaultModel         = "gpt-5"
	DefaultCommandPrefix = "/"
	DefaultTimeout       = 60
	DefaultColumns       = 80
)

const (
	APP_DIR  = "cli-chat"
	FILENAME = "settings.json"
)

type Config struct {
	Client      *client.Client
	AppSettings *Settings
}

type Provider struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	BaseURL string `json:"baseurl"`
}

type Settings struct {
	Provider      Provider `json:"provider"`
	Model         string   `json:"model"`
	CommandPrefix string   `json:"command_prefix"`
	Timeout       int      `json:"timeout"`
	Columns       int      `json:"columns"`
}

var DefaultProvider = Provider{
	Name:    "openai",
	Key:     "OPENAI_API_KEY",
	BaseURL: "https://api.openai.com/v1/",
}

var DefaultSettings = Settings{
	Provider:      DefaultProvider,
	Model:         DefaultModel,
	CommandPrefix: DefaultCommandPrefix,
	Timeout:       DefaultTimeout,
	Columns:       DefaultColumns,
}

func NewConfig(s *Settings) (*Config, error) {
	cfg := &Config{
		Client:      client.New(time.Duration(s.Timeout) * time.Second),
		AppSettings: s,
	}
	return cfg, nil
}

func NewDefaultSettings() Settings {
	return DefaultSettings
}
