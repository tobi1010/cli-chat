package commands

import (
	"cli-chat/session"
	"fmt"
)

func CommandSwitchModel(s *session.Session, args []string) error {
	s.Provider.Model = args[0]
	if err := s.Save(s.Paths.SessionPath); err != nil {
		return fmt.Errorf("saving session: %w", err)
	}
	return nil
}
