package api

import (
	"cli-chat/internal/anthropic"
	"cli-chat/internal/apitypes"
	"cli-chat/internal/client"
	"cli-chat/internal/openai"
	"cli-chat/providers"
	"context"
	"fmt"
)

func GetModels(ctx context.Context, client *client.Client, provider providers.Provider) ([]apitypes.Model, error) {
	switch provider.Name {
	case "openai":
		return openai.GetModels(ctx, client, provider)
	case "anthropic":
		return anthropic.GetModels(ctx, client, provider)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", provider.Name)
	}
}
func CreateStreamResponse(ctx context.Context, client *client.Client, provider providers.Provider, input string) (<-chan string, <-chan apitypes.Response, <-chan error, error) {
	switch provider.Name {
	case "openai":
		return openai.CreateStreamResponse(ctx, client, provider, input)
	case "anthropic":
		return anthropic.CreateStreamResponse(ctx, client, provider, input)
	default:
		return nil, nil, nil, fmt.Errorf("unsupported provider: %s", provider.Name)
	}
}
