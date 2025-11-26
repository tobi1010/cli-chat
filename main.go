package main

import (
	"cli-chat/config"
	"log"
	"os"
)

const CWD = "."

func main() {
	s, err := config.ReadSettings(CWD)
	config.PrintSettings(CWD)
	cfg, err := config.NewConfig(s)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = startRepl(&cfg)
	if err != nil {
		log.Printf("%v\n", err)
	}
}
