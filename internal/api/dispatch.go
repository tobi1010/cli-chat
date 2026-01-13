package api

import (
	"terminal-chat/chat"
	"terminal-chat/internal/anthropic"
	"terminal-chat/internal/apitypes"
	"terminal-chat/internal/client"
	"terminal-chat/internal/openai"
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
func CreateStreamResponse(ctx context.Context, client *client.Client, providerName string, modelID, systemPrompt string, messages []chat.Message) (<-chan string, <-chan apitypes.Response, <-chan error, error) {
	switch providerName {
	case "openai":
		return openai.CreateStreamResponse(ctx, client, providerName, modelID, systemPrompt, messages)
	case "anthropic":
		return anthropic.CreateStreamResponse(ctx, client, providerName, modelID, systemPrompt, messages)
	default:
		return nil, nil, nil, fmt.Errorf("unsupported provider: %s", providerName)
	}
}
