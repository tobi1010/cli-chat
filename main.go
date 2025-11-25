package main

import (
	"cli-chat/settings"
	"log"
	"os"
)

const CWD = "."

func main() {
	s, err := settings.Read(CWD)
	settings.Print(CWD)
	cfg, err := newConfig(s)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	startRepl(&cfg)
}
