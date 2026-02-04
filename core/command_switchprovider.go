package core

import (
	"fmt"
	"terminal-chat/providers"
)

func CommandSwitchProvider(s *Session, args []string) error {
	if len(args) < 1 {
		fmt.Printf("usage: /switch-provider <name>\n")
		fmt.Println("switching to default provider...")
		s.ProviderName, s.ModelID = providers.DefaultNameAndModelID()
		return s.SessionSave(s.Paths.SessionPath)
	}

	def, ok := providers.Get(args[0])
	if !ok {
		fmt.Printf("unknown provider: %s\n", args[0])
		s.ProviderName, s.ModelID = providers.DefaultNameAndModelID()
		return s.SessionSave(s.Paths.SessionPath)
	}

	s.ProviderName = def.Name
	s.ModelID = def.DefaultModel

	return s.SessionSave(s.Paths.SessionPath)
}
