package commands

import (
	"cli-chat/providers"
	"cli-chat/session"
	"fmt"
)

func CommandSwitchProvider(s *session.Session, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: /switch-provider <name>")
	}

	name := args[0]
	prv, ok := providers.Get(name, "")
	if !ok {
		return fmt.Errorf("unknown provider: %s", name)
	}

	s.Provider = prv

	if err := s.Save(s.Paths.SessionPath); err != nil {
		return fmt.Errorf("saving session: %w", err)
	}

	return nil
}
