package main

import (
	"bufio"
	"cli-chat/internal/api"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name     string
	args     []string
	callback func(*Config, string) error
}

var commands = map[string]cliCommand{
	"exit": {
		name:     "exit",
		callback: commandExit,
	},
}

func startRepl(cfg *Config) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("llm>")
		scanner.Scan()

		text := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(text, cfg.Settings.CommandPrefix) {
			// command
			fmt.Println("Command!")
			lowerText := strings.ToLower(text)
			tokens := strings.Fields(lowerText)
			for _, token := range tokens {

				if command, found := commands[token]; found {
					err := command.callback(cfg, "")
					if err != nil {
						return fmt.Errorf("Error executing command: %w", err)
					}
				}
			}

		} else {
			err := api.StreamResponse(cfg.Client, cfg.Settings.Model, text)
			if err != nil {
				return fmt.Errorf("Error creating response: %w", err)
			}
		}
	}
}
