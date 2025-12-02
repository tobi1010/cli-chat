package main

import (
	"cli-chat/config"
	"log"
	"os"
)

func main() {
	err := config.InitDefaultSettings()
	if err != nil {
		log.Println("initializing default settings: %v", err)
		os.Exit(1)
	}

	var s config.Settings
	err = config.ReadSettings(&s)
	if err != nil {
		log.Println("reading settings: %v", err)
		os.Exit(1)
	}

	err = config.PrintSettings()
	if err != nil {
		log.Println("printing settings: %v", err)
		os.Exit(1)
	}

	cfg, err := config.NewConfig(&s)
	if err != nil {
		log.Println("creating config: %v", err)
		os.Exit(1)
	}

	err = startRepl(cfg)
	if err != nil {
		log.Println("repl: %v", err)
		os.Exit(1)
	}
}
