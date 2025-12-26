package api

import (
	// "cli-chat/internal/anthropic"
	"cli-chat/internal/llm"
	"cli-chat/internal/openai"
	"cli-chat/session"
	"context"
	"fmt"
)

var backends = map[string]llm.Backend{
	"openai": openai.Backend{},
	// "anthropic": anthropic.Backend{},
}

func backendFor(s *session.Session) (llm.Backend, error) {
	b, ok := backends[s.Provider.Name]
	if !ok {
		return nil, fmt.Errorf("unsupported provider: %s", s.Provider.Name)
	}
	return b, nil
}

func GetModels(ctx context.Context, s *session.Session) ([]llm.Model, error) {
	b, err := backendFor(s)
	if err != nil {
		return nil, err
	}
	return b.GetModels(ctx, s)
}

func CreateStreamResponse(
	ctx context.Context,
	s *session.Session,
	input string,
) (<-chan string, <-chan llm.Response, <-chan error, error) {
	b, err := backendFor(s)
	if err != nil {
		return nil, nil, nil, err
	}
	return b.CreateStreamResponse(ctx, s, input)
}
