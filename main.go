package main

import (
	// "cli-chat/commands"
	"cli-chat/paths"
	"cli-chat/session"
	"errors"
	"log"
	"os"
)

func main() {
	pths, err := paths.ResolvePaths()
	if err != nil {
		log.Printf("error resolving paths:")
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}

	s, err := session.Open(pths)
	if err != nil {
		log.Printf("error oopening session %s:", paths.SessionPath)
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}
	err = s.Save(pths.SessionPath)
	if err != nil {
		log.Printf("error saving session to %s:", paths.SessionPath)
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}

	// if err := commands.CommandPrintSettings(s, []string{}); err != nil {
	// 	log.Printf("error printing settings:")
	// 	unwrapAndPrintErrors(err)
	// 	os.Exit(1)
	// }

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
