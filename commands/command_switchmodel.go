package commands

import (
	"cli-chat/providers"
	"cli-chat/session"
	"fmt"
)

func CommandSwitchModel(s *session.Session, args []string) error {
	if len(args) < 1 {
		fmt.Printf("usage: /switch-model <name>\n")
		return nil
	}
	prov, err := providers.New(*s.Cache, s.Provider.Name, args[0])
	if err != nil {
		return fmt.Errorf("unknown model %s for provider %s: %w", args[0], s.Provider.Name, err)
	}
	s.Provider = prov

	if err := s.Save(s.Paths.SessionPath); err != nil {
		return fmt.Errorf("saving session: %w", err)
	}
	return nil
}
