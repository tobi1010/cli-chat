package config

import (
	"cli-chat/internal/client"
	"time"
)

type Config struct {
	SettingsPath string
	Client       *client.Client
	AppSettings  *Settings
}

type Provider struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	BaseURL string `json:"baseurl"`
}

type Settings struct {
	Provider      Provider `json:"provider"`
	Model         string   `json:"model"`
	LastChatID    string   `json:"last_chat_id"`
	CommandPrefix string   `json:"command_prefix"`
	Timeout       int      `json:"timeout"`
	Columns       int      `json:"columns"`
}

func NewDefaultSettings() Settings {
	return Settings{
		Provider: Provider{
			Name:    "openai",
			Key:     "OPENAI_API_KEY",
			BaseURL: "https://api.openai.com/v1/",
		},
		Model:         "gpt-5",
		CommandPrefix: "/",
		Timeout:       60,
		Columns:       80,
	}
}

func New(settingsPath string, s *Settings) (*Config, error) {
	cfg := &Config{
		SettingsPath: settingsPath,
		Client:       client.New(time.Duration(s.Timeout) * time.Second),
		AppSettings:  s,
	}
	return cfg, nil
}
