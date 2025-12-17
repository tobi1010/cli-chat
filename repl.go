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
		fmt.Printf("%s>", s.Provider.Model)
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
				if err := command.Callback(s, tokens[1:]); err != nil {
					return fmt.Errorf("Error executing command: %w", err)
				}
			} else {
				fmt.Printf("unknown command: %sp\n", tokens[0])
			}

		} else {
			ctx := context.Background()
			stream, fullCh, errCh, err := api.CreateStreamResponse(ctx, s, text)
			if err != nil {
				return fmt.Errorf("recieving stream : %w", err)
			}
			s.Chat.AddMessage("user", text)
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
					outputText := ""
					role := ""
					if ok {
						for i := range fullResponse.Output {
							if fullResponse.Output[i].Role != "" {
								role = fullResponse.Output[i].Role
							}
							for j := range fullResponse.Output[i].Content {
								outputText = outputText + fullResponse.Output[i].Content[j].Text
							}
						}

						s.Chat.AddMessage(role, outputText)
						fmt.Printf("%v", s.Chat.Conversation)
						_ = fullResponse
					}
				}
			}
			fmt.Println()
		}
	}
}
