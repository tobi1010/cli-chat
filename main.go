package main

import (
	"cli-chat/config"
	"log"
	"os"
)

func main() {
	settings, settingsPath, err := config.EnsureSettings()
	if err != nil {
		log.Printf("initializing default settings: %v", err)
		os.Exit(1)
	}

	cfg, err := config.New(settingsPath, &settings)
	if err != nil {
		log.Printf("creating config: %v", err)
		os.Exit(1)
	}

	err = startRepl(cfg)
	if err != nil {
		log.Printf("repl: %v", err)
		os.Exit(1)
	}
}
