package api

import (
	"cli-chat/internal/anthropic"
	"cli-chat/internal/apitypes"
	"cli-chat/internal/openai"
	"cli-chat/session"
	"context"
	"fmt"
)

func GetModels(ctx context.Context, s *session.Session) ([]apitypes.Model, error) {
	switch s.Provider.Name {
	case "openai":
		return openai.GetModels(ctx, s)
	case "anthropic":
		return anthropic.GetModels(ctx, s)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", s.Provider.Name)
	}
}
func CreateStreamResponse(ctx context.Context, s *session.Session, input string) (<-chan string, <-chan apitypes.Response, <-chan error, error) {
	return nil, nil, nil, nil
}
