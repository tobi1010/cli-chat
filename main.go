package main

import (
	"cli-chat/config"
	"log"
	"os"
)

func main() {
	err := config.InitDefaultSettings()
	if err != nil {
		log.Printf("initializing default settings: %v", err)
		os.Exit(1)
	}

	var s config.Settings
	err = config.ReadSettings(&s)
	if err != nil {
		log.Printf("reading settings: %v", err)
		os.Exit(1)
	}

	if err != nil {
		log.Printf("printing settings: %v", err)
		os.Exit(1)
	}

	cfg, err := config.NewConfig(&s)
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
