package main

import (
	"cli-chat/paths"
	"cli-chat/session"
	"log"
	"os"
)

func main() {
	sessionPath, err := paths.SessionPath()
	s, err := session.LoadOrCreate(sessionPath)
	if err != nil {
		log.Printf("creating session: %v", err)
		os.Exit(1)
	}

	err = startRepl(s)
	if err != nil {
		log.Printf("repl: %v", err)
		os.Exit(1)
	}
}
