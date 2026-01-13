package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"terminal-chat/chat"
	"terminal-chat/commands"
	"terminal-chat/internal/api"
	"terminal-chat/internal/apitypes"
	"terminal-chat/session"
	"terminal-chat/spinner"
	"terminal-chat/util"
)

func startRepl(s *session.Session) error {
	cmds := commands.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)
	util.ClearScreen()
	util.PrintTitle()
	fmt.Println(strings.Repeat("-", s.AppSettings.Columns))
	fmt.Printf("%shelp for command list\n", s.AppSettings.CommandPrefix)
	fmt.Println(strings.Repeat("-", s.AppSettings.Columns))
	for {
		fmt.Printf("%s>", s.ModelLabel)
		text, ok, err := readInput(scanner)
		if err != nil {
			return fmt.Errorf("reading input: %w", err)
		}
		if !ok {
			return nil //EOF
		}
		if text == "" {
			continue
		}
		handled, err := handleCommand(s, cmds, text)
		if err != nil {
			if errors.Is(err, commands.ErrExitRequested) {
				return nil
			}
			fmt.Println(err.Error())
		}
		if handled {
			continue
		}
		if err := handleChat(s, text); err != nil {
			return fmt.Errorf("handling chat: %w", err)
		}

	}
}

func readInput(scanner *bufio.Scanner) (string, bool, error) {
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", false, fmt.Errorf("reading stdin: %w", err)
		}
		return "", false, nil //EOF
	}
	return strings.TrimSpace(scanner.Text()), true, nil
}

func handleCommand(s *session.Session, cmds map[string]commands.CliCommand, text string) (bool, error) {
	if !strings.HasPrefix(text, s.AppSettings.CommandPrefix) {
		return false, nil
	}
	// command
	lowerText := strings.ToLower(strings.TrimPrefix(text, s.AppSettings.CommandPrefix))
	tokens := strings.Fields(lowerText)
	if len(tokens) == 0 {
		return true, nil
	}

	if command, found := cmds[tokens[0]]; found {
		if err := command.Callback(s, tokens[1:]); err != nil {
			if errors.Is(err, commands.ErrExitRequested) {
				return true, err
			}
			return true, fmt.Errorf("Error executing command: %w", err)
		}
	} else {
		fmt.Printf("unknown command: %s\n", tokens[0])
		return true, nil
	}
	return true, nil
}

func handleChat(s *session.Session, text string) error {
	fmt.Println(strings.Repeat("-", s.AppSettings.Columns))
	ctx := context.Background()
	s.Chat.AddMessage(chat.Message{Role: "user", Content: text})
	stream, fullCh, errCh, err := api.CreateStreamResponse(ctx, s.Client, s.ProviderName, s.ModelID, s.SystemPrompt, s.Chat.Conversation)
	if err != nil {
		return fmt.Errorf("receiving stream : %w", err)
	}

	fullResponse, err := consumeStream(stream, fullCh, errCh, s.AppSettings.Columns)
	if err != nil {
		return fmt.Errorf("consuming stream: %w", err)
	}
	msg := extractMessage(fullResponse)

	s.Chat.AddMessage(msg)
	if err := s.UpdateChat(); err != nil {
		return fmt.Errorf("updating chat: %w", err)
	}
	fmt.Printf("\n        response by: %s\n", fullResponse.Model)

	fmt.Println(strings.Repeat("-", s.AppSettings.Columns))
	return nil
}

func consumeStream(stream <-chan string, fullCh <-chan apitypes.Response, errCh <-chan error, printColumns int) (apitypes.Response, error) {
	lineLength := 0
	streamOpen := true
	stopSpinner := spinner.Spin()
	for streamOpen {
		select {
		case delta, ok := <-stream:
			stopSpinner()
			if !ok {
				streamOpen = false
				break
			}
			fmt.Print(delta)
			lineLength += len(delta)
			if lineLength > printColumns {
				fmt.Print("\n")
				lineLength = 0
			}

		case err, ok := <-errCh:
			if !ok {
				errCh = nil
				break
			}
			if err != nil {
				return apitypes.Response{}, fmt.Errorf("stream error: %w", err)
			}
		case fullResponse, ok := <-fullCh:
			if !ok {
				return apitypes.Response{}, fmt.Errorf("stream ended without full response")
			}
			return fullResponse, nil
		}
	}

	return apitypes.Response{}, fmt.Errorf("stream closed without full response")
}

func extractMessage(fullResponse apitypes.Response) chat.Message {
	var b strings.Builder

	for _, out := range fullResponse.Output {
		// Only take assistant "message" outputs (skip tool calls etc).
		if out.Type != "message" {
			continue
		}
		if out.Role != "assistant" {
			continue
		}

		for _, part := range out.Content {
			// Only take text parts.
			if part.Type == "output_text" || part.Type == "text" {
				b.WriteString(part.Text)
			}
		}
	}

	return chat.Message{
		Role:    "assistant",
		Content: b.String(),
	}
}
