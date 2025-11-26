package config

import (
	"cli-chat/internal/client"
	"time"
)

type Config struct {
	Client      *client.Client
	AppSettings Settings
}

func NewConfig(s Settings) (Config, error) {
	cfg := Config{
		Client:      client.New(time.Duration(s.Timeout) * time.Second),
		AppSettings: s,
	}
	return cfg, nil
}
