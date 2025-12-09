package main

import (
	"bufio"
	"cli-chat/commands"
	"cli-chat/internal/api"
	"cli-chat/session"
	"context"
	"fmt"
	"os"
	"strings"
)

func startRepl(s *session.Session) error {
	cmds := commands.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s>", s.Model)
		scanner.Scan()

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		if strings.HasPrefix(text, s.AppSettings.CommandPrefix) {
			// command
			lowerText := strings.ToLower(strings.TrimPrefix(text, s.AppSettings.CommandPrefix))
			tokens := strings.Fields(lowerText)

			if command, found := cmds[tokens[0]]; found {
				err := command.Callback(s, tokens[1:])
				if err != nil {
					return fmt.Errorf("Error executing command: %w", err)
				}
			}

		} else {
			ctx := context.Background()
			stream, fullCh, errCh, err := api.CreateStreamResponse(ctx, s, text)
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
					if lineLength > s.AppSettings.Columns {
						fmt.Print("\n")
						lineLength = 0
					}

				case err, ok := <-errCh:
					if ok && err != nil {
						return fmt.Errorf("stream error: %w", err)
					}
				case fullResponse, ok := <-fullCh:
					if ok {
						_ = fullResponse
					}
				}
			}
			fmt.Println()
		}
	}
}
