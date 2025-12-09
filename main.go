package main

import (
	"cli-chat/session"
	"log"
	"os"
)

func main() {
	s, err := session.LoadOrCreate()
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
