package core

import (
	"fmt"
)

func CommandSetSystemPrompt(s *Session, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: %sset-system-prompt <prompt>", s.AppSettings.CommandPrefix)
	}
	s.SystemPrompt = args[0]
	return nil
}
