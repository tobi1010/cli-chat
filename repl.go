package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name     string
	args     []string
	callback func(*Config, string) error
}

func commandExit(*Config, string) error {
	fmt.Println("EXIT")
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:     "exit",
			callback: commandExit,
		},
	}
}

func startRepl(cfg *Config) error {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("llm>")
		scanner.Scan()

		text := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(text, cfg.Command_delimiter) {
			// command
			lowerText := strings.ToLower(text)
			tokens := strings.Fields(lowerText)
			for _, token := range tokens {

				if command, found := commands[token]; found {
					err := command.callback(cfg, "")
					if err != nil {
						return fmt.Errorf("Error executing command: %w", err)
					}
				} else {
					// send text to llm
				}
			}
		}
	}
}
