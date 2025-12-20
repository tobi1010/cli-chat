package main

import (
	"bufio"
	"cli-chat/commands"
	"cli-chat/internal/api"
	"cli-chat/session"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
)

func startRepl(s *session.Session) error {
	cmds := commands.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s>", s.Provider.Model)
		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				return fmt.Errorf("reading stdin: %w", err)
			}
			return nil //EOF
		}

		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			continue
		}

		if strings.HasPrefix(text, s.AppSettings.CommandPrefix) {
			// command
			lowerText := strings.ToLower(strings.TrimPrefix(text, s.AppSettings.CommandPrefix))
			tokens := strings.Fields(lowerText)
			if len(tokens) == 0 {
				continue
			}

			if command, found := cmds[tokens[0]]; found {
				if err := command.Callback(s, tokens[1:]); err != nil {
					if errors.Is(err, commands.ErrExitRequested) {
						return nil
					}
					return fmt.Errorf("Error executing command: %w", err)
				}
			} else {
				fmt.Printf("unknown command: %s\n", tokens[0])
			}

		} else {
			ctx := context.Background()
			s.Chat.AddMessage("user", text)
			stream, fullCh, errCh, err := api.CreateStreamResponse(ctx, s, fmt.Sprintf("%v", s.Chat.Conversation))
			if err != nil {
				return fmt.Errorf("receiving stream : %w", err)
			}
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
					lineLength += len(delta)
					if lineLength > s.AppSettings.Columns {
						fmt.Print("\n")
						lineLength = 0
					}

				case err, ok := <-errCh:
					if !ok {
						errCh = nil
						break
					}
					if ok && err != nil {
						return fmt.Errorf("stream error: %w", err)
					}
				case fullResponse, ok := <-fullCh:
					var outputText strings.Builder
					role := ""
					if ok {
						for i := range fullResponse.Output {
							if fullResponse.Output[i].Role != "" {
								role = fullResponse.Output[i].Role
							}
							for j := range fullResponse.Output[i].Content {
								outputText.WriteString(fullResponse.Output[i].Content[j].Text)
							}
						}

						s.Chat.AddMessage(role, outputText.String())
						if err := s.UpdateChat(); err != nil {
							return fmt.Errorf("updating chat: %w", err)
						}
						fmt.Printf("\n        response by: %s\n", fullResponse.Model)
					}
				}
			}
			fmt.Println()
		}
	}
}
