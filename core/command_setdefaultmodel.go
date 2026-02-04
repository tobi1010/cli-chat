package core

import (
	"fmt"
)

func CommandSetDefautlModel(s *Session, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("usage: %sset-default-model <model>", s.AppSettings.CommandPrefix)
	}
	model, ok := s.Cache.Get(s.AppSettings.DefaultProvider, args[0])
	if !ok {
		return fmt.Errorf("unknown model %q for provider %q", args[0], s.AppSettings.DefaultProvider)
	}
	s.AppSettings.DefaultModel = model.ID
	return s.SessionSave(s.Paths.SettingsPath)
}
