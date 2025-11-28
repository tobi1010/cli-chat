package main

import (
	"bufio"
	"cli-chat/config"
	"cli-chat/internal/api"
	"context"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name     string
	args     []string
	callback func(*config.Config, []string) error
}

var commands = map[string]cliCommand{
	"exit": {
		name:     "exit",
		callback: commandExit,
	},
	"print-settings": {
		name:     "print-settings",
		callback: commandPrintSettings,
	},
	"set-prefix": {
		name:     "set-prefix",
		callback: commandSetPrefix,
	},
	"set-columns": {
		name:     "set-columns",
		callback: commandSetColumns,
	},
	"list-chats": {
		name:     "list-chats",
		callback: commandListChats,
	},
	"list-models": {
		name:     "list-models",
		callback: commandListModels,
	},
	"list-providers": {
		name:     "list-prividers",
		callback: commandListProviders,
	},
	"usage": {
		name:     "usage",
		callback: commandUsage,
	},
}

func startRepl(cfg *config.Config) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("llm>")
		scanner.Scan()

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		if strings.HasPrefix(text, cfg.AppSettings.CommandPrefix) {
			// command
			lowerText := strings.ToLower(strings.TrimPrefix(text, cfg.AppSettings.CommandPrefix))
			tokens := strings.Fields(lowerText)

			if command, found := commands[tokens[0]]; found {
				err := command.callback(cfg, tokens[1:])
				if err != nil {
					return fmt.Errorf("Error executing command: %w", err)
				}
			}

		} else {
			ctx := context.Background()
			stream, err := api.CreateStreamResponse(ctx, cfg, text)
			if err != nil {
				return fmt.Errorf("recieving stream : %w", err)
			}
			var acc strings.Builder
			lineLength := 0
			for delta := range stream {
				fmt.Print(delta)
				len, err := acc.WriteString(delta)
				if err != nil {
					return fmt.Errorf("building response string: %w", err)
				}
				lineLength += len
				if lineLength > cfg.AppSettings.Columns {
					fmt.Print("\n")
					lineLength = 0
				}
			}
			fmt.Println()
		}
	}
}
