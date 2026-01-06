package anthropic

import (
	"cli-chat/internal/apitypes"
	"cli-chat/internal/client"
	"cli-chat/providers"
	"context"
)

func CreateStreamResponse(ctx context.Context, client *client.Client, provider providers.Provider, input string) (<-chan string, <-chan apitypes.Response, <-chan error, error) {
	return nil, nil, nil, nil
}
