package anthropic

import (
	"terminal-chat/chat"
	"terminal-chat/internal/apitypes"
	"terminal-chat/internal/client"
	"terminal-chat/providers"
	"context"
	"fmt"
)

func CreateStreamResponse(ctx context.Context, client *client.Client, providerName, modelID string, systemPrompt string, messages []chat.Message) (<-chan string, <-chan apitypes.Response, <-chan error, error) {
	_, ok := providers.Get(providerName)
	if !ok {
		return nil, nil, nil, fmt.Errorf("unknown provider %q", providerName)
	}

	return nil, nil, nil, nil
}
