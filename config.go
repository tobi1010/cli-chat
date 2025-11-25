package main

import (
	"cli-chat/internal/client"
	"cli-chat/settings"
	"time"
)

type Config struct {
	Client      *client.Client
	AppSettings settings.Settings
}

func newConfig(s settings.Settings) (Config, error) {
	cfg := Config{
		Client:      client.New(time.Duration(s.Timeout) * time.Second),
		AppSettings: s,
	}
	return cfg, nil
}
