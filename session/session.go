package session

import (
	"cli-chat/internal/api"
	"cli-chat/internal/apitypes"
	"cli-chat/providers"
	"context"
	"fmt"
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
	provider := providers.GetByName(providerName)
	models, err := api.GetModels(context.Background(), s.Client, provider)
	if err != nil {
		return nil, fmt.Errorf("fetching models for %q", s.Provider.Name)
	}
	return models, nil
}
