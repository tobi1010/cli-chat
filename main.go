package main

import (
	"cli-chat/config"
	"log"
)

func main() {
	err := config.InitDefaultSettings()
	if err != nil {
		log.Fatalf("initializing default settings: %v", err)
	}

	var s config.Settings
	err = config.ReadSettings(&s)
	if err != nil {
		log.Fatalf("reading settings: %v", err)
	}

	err = config.PrintSettings()
	if err != nil {
		log.Fatalf("printing settings: %v", err)
	}

	cfg, err := config.NewConfig(&s)
	if err != nil {
		log.Fatalf("creating config: %v", err)
	}

	err = startRepl(cfg)
	if err != nil {
		log.Fatalf("repl: %v", err)
	}
}
