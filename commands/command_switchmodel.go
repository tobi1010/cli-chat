package commands

import (
	"context"
	"fmt"
	"terminal-chat/session"
	"time"
)

func CommandSwitchModel(s *session.Session, args []string) error {
	if len(args) < 1 {
		fmt.Printf("usage: /switch-model <name>\n")
		return nil
	}
	model, ok := s.Cache.Get(s.ProviderName, args[0])
	if !ok {
		err := s.Cache.EnsureFresh(context.Background(), s.Paths.CachePath, s.ProviderName, time.Duration(s.AppSettings.TTL)*time.Second, s)
		if err != nil {
			return fmt.Errorf("refreshing cache: %w", err)
		}
		model, ok = s.Cache.Get(s.ProviderName, args[0])
		if !ok {
			return fmt.Errorf("unknown model: %q", args[0])
		}
	}
	s.ModelID = model.ID

	if err := s.Save(s.Paths.SessionPath); err != nil {
		return fmt.Errorf("saving session: %w", err)
	}
	return nil
}
