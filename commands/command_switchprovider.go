package commands

import (
	"cli-chat/providers"
	"cli-chat/session"
	"fmt"
)

func CommandSwitchProvider(s *session.Session, args []string) error {
	if len(args) < 1 {
		fmt.Printf("usage: /switch-provider <name>\n")
		return nil
	}

	if args[0] == "" {
		fmt.Println("switching to default provider...")

		prov, err := providers.NewDefault(*s.Cache)
		if err != nil {
			return fmt.Errorf("switching to default provider")
		}
		s.Provider = prov

	}
	name := args[0]
	prov, err := providers.New(*s.Cache, args[0], "")
	if err != nil {
		return fmt.Errorf("unknown provider: %s", name)
	}
	s.Provider = prov

	if err := s.Save(s.Paths.SessionPath); err != nil {
		return fmt.Errorf("saving session: %w", err)
	}

	return nil
}
