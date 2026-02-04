package main

import (
	// "terminal-chat/commands"
	"errors"
	"flag"
	"log"
	"os"
	"terminal-chat/core"
	"terminal-chat/debug"
	"terminal-chat/paths"
)

func main() {
	debugFlag := flag.Bool("debug", false, "enable debug dumps")
	flag.Parse()
	debug.Set(*debugFlag)
	CommandRegistry, err := core.NewRegistry()
	CommandMeta := core.NewCommandMeta()
	_ = CommandMeta

	if err != nil {
		log.Printf("error building command Registry:")
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}

	pths, err := paths.ResolvePaths()
	if err != nil {
		log.Printf("error resolving paths:")
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}

	s, err := core.SessionOpen(pths)
	if err != nil {
		log.Printf("error oopening session %s:", pths.SessionPath)
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}
	err = s.SessionSave(pths.SessionPath)
	if err != nil {
		log.Printf("error saving session to %s:", pths.SessionPath)
		unwrapAndPrintErrors(err)
		os.Exit(1)
	}

	if err := startRepl(s, CommandRegistry); err != nil {
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
