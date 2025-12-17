package main

import (
	"cli-chat/commands"
	"cli-chat/paths"
	"cli-chat/session"
	"log"
	"os"
)

func main() {
	sessionPath, err := paths.SessionPath()
	s, err := session.LoadOrCreate(sessionPath)
	commands.CommandPrintSettings(s, []string{})
	if err != nil {
		log.Printf("creating session: %v", err)
		os.Exit(1)
	}

	if err = startRepl(s); err != nil {
		log.Printf("repl: %v", err)
		os.Exit(1)
	}
}
