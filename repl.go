package main

import (
	"bufio"
	"cli-chat/commands"
	"cli-chat/config"
	"cli-chat/internal/api"
	"context"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config.Config) error {
	cmds := commands.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s>", cfg.AppSettings.Model)
		scanner.Scan()

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		if strings.HasPrefix(text, cfg.AppSettings.CommandPrefix) {
			// command
			lowerText := strings.ToLower(strings.TrimPrefix(text, cfg.AppSettings.CommandPrefix))
			tokens := strings.Fields(lowerText)

			if command, found := cmds[tokens[0]]; found {
				err := command.Callback(cfg, tokens[1:])
				if err != nil {
					return fmt.Errorf("Error executing command: %w", err)
				}
			}

		} else {
			ctx := context.Background()
			stream, fullCh, errCh, err := api.CreateStreamResponse(ctx, cfg, text)
			if err != nil {
				return fmt.Errorf("recieving stream : %w", err)
			}
			var acc strings.Builder
			lineLength := 0
			streamOpen := true
			for streamOpen {
				select {
				case delta, ok := <-stream:
					if !ok {
						streamOpen = false
						break
					}
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

				case err, ok := <-errCh:
					if ok && err != nil {
						return fmt.Errorf("stream error: %w", err)
					}
				case fullResponse, ok := <-fullCh:
					if ok {
						fmt.Println()
						fmt.Println(fullResponse)
						fmt.Println()
					}
				}
			}
			fmt.Println()
		}
	}
}
