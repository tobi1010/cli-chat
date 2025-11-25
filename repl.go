package main

import (
	"bufio"
	"cli-chat/internal/api"
	"context"
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

		if strings.HasPrefix(text, cfg.AppSettings.CommandPrefix) {
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
			ctx := context.Background()
			deltas, fullCh, errCh, cancel, err := api.StreamDeltas(ctx, cfg.Client, cfg.AppSettings.Model, text)
			if err != nil {
				return fmt.Errorf("streaming deltas: %w", err)
			}
			defer cancel()
			var buf strings.Builder
			for d := range deltas {
				fmt.Print(d)
				buf.WriteString(d)
			}
			accText := buf.String()

			select {
			case finalRes := <-fullCh:
				fmt.Printf("ID: %s\n Model: %s", finalRes.Response.ID, finalRes.Response.Model)

			case e := <-errCh:
				if e != nil {
					fmt.Println("stream error:", e)
				}
			default:
			}
			_ = accText

			if err != nil {
				return fmt.Errorf("Error creating response: %w", err)
			}
		}
	}
}
