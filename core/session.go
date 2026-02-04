package core

import (
	"context"
	"fmt"
	"terminal-chat/internal/api"
	"terminal-chat/internal/apitypes"
	"terminal-chat/providers"
)

func (s *Session) UpdateChat() error {
	if s.DB == nil {
		return fmt.Errorf("db is nil")
	}
	if s.Chat == nil {
		return fmt.Errorf("chat is nil")
	}
	if err := s.Chat.Save(s.Paths.ChatsDir); err != nil {
		return fmt.Errorf("writing chat atomically: %w", err)
	}
	s.DB.Touch(s.Chat.ID, s.Chat.UpdatedAt)
	if err := s.DB.Save(s.Paths.IndexPath); err != nil {
		return fmt.Errorf("saving db: %w", err)
	}
	return nil
}

func (s *Session) FetchModels(ctx context.Context, providerName string) ([]apitypes.Model, error) {
	provider, ok := providers.Get(providerName)
	if !ok {
		return nil, fmt.Errorf("unknown provider: %q", provider)
	}
	models, err := api.GetModels(context.Background(), s.Client, provider.ID)
	if err != nil {
		return nil, fmt.Errorf("fetching models for %q", s.ProviderName)
	}
	return models, nil
}
