package main

import (
	"cli-chat/commands"
	"cli-chat/paths"
	"cli-chat/session"
	"errors"
	"log"
	"os"
)

func main() {
	sessionPath, err := paths.SessionPath()
	if err != nil {
		log.Printf("error resolving session path:")
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}

	s, err := session.LoadOrCreate(sessionPath)
	if err != nil {
		log.Printf("error LoadOrCreate %s:", sessionPath)
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}
	_ = s.Save(sessionPath)

	if err := commands.CommandPrintSettings(s, []string{}); err != nil {
		log.Printf("error printing settings:")
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}

	if err := startRepl(s); err != nil {
		log.Printf("error in repl:")
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func unwrapAndPrintErrors(err error) {
	for e := err; e != nil; e = errors.Unwrap(e) {
		log.Printf("  -> %v", e)
	}
}
