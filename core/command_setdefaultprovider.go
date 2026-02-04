package core

import (
	"fmt"
	"terminal-chat/providers"
)

func CommandSetDefautlProvider(s *Session, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: %sset-default-provider <provider>", s.AppSettings.CommandPrefix)
	}
	provider, ok := providers.Get(args[0])
	if !ok {
		return fmt.Errorf("unknown provider: %s", args[0])
	}
	s.AppSettings.DefaultProvider = provider.ID
	return s.SessionSave(s.Paths.SettingsPath)
}
