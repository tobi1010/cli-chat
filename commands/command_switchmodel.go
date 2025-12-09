package commands

import (
	"cli-chat/session"
	"fmt"
)

func CommandSwitchModel(s *session.Session, args []string) error {
	s.Model = args[0]
	err := s.Chat.Write()
	if err != nil {
		return fmt.Errorf("writing settings.json: %w", err)
	}
	return nil
}
