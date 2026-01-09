package commands

import (
	"cli-chat/providers"
	"cli-chat/session"
	"fmt"
)

func CommandSwitchProvider(s *session.Session, args []string) error {
	if len(args) < 1 {
		fmt.Printf("usage: /switch-provider <name>\n")
		fmt.Println("switching to default provider...")
		s.ProviderName, s.ModelID = providers.DefaultNameAndModelID()
		return s.Save(s.Paths.SessionPath)
	}

	def, ok := providers.Get(args[0])
	if !ok {
		fmt.Printf("unknown provider: %s\n", args[0])
		s.ProviderName, s.ModelID = providers.DefaultNameAndModelID()
		return s.Save(s.Paths.SessionPath)
	}

	s.ProviderName = def.Name
	s.ModelID = def.DefaultModel

	return s.Save(s.Paths.SessionPath)
}
