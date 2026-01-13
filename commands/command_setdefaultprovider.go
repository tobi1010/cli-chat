package commands

import (
	"fmt"
	"terminal-chat/providers"
	"terminal-chat/session"
)

func CommandSetDefautlProvider(s *session.Session, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: %sset-default-provider <provider>", s.AppSettings.CommandPrefix)
	}
	provider, ok := providers.Get(args[0])
	if !ok {
		return fmt.Errorf("unknown provider: %s", args[0])
	}
	s.AppSettings.DefaultProvider = provider.ID
	return s.Save(s.Paths.SettingsPath)
}
