package api

import (
	"cli-chat/internal/anthropic"
	"cli-chat/internal/apitypes"
	"cli-chat/internal/client"
	"cli-chat/internal/openai"
	"context"
	"fmt"
)

func GetModels(ctx context.Context, client *client.Client, providerName string) ([]apitypes.Model, error) {
	switch providerName {
	case "openai":
		return openai.GetModels(ctx, client, providerName)
	case "anthropic":
		return anthropic.GetModels(ctx, client, providerName)
	default:
		return nil, fmt.Errorf("unsupported provider: %s", providerName)
	}
}
func CreateStreamResponse(ctx context.Context, client *client.Client, providerName string, modelID, input string) (<-chan string, <-chan apitypes.Response, <-chan error, error) {
	switch providerName {
	case "openai":
		return openai.CreateStreamResponse(ctx, client, providerName, modelID, input)
	case "anthropic":
		return anthropic.CreateStreamResponse(ctx, client, providerName, modelID, input)
	default:
		return nil, nil, nil, fmt.Errorf("unsupported provider: %s", providerName)
	}
}
