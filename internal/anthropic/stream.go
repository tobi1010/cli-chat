package anthropic

import (
	"cli-chat/internal/apitypes"
	"cli-chat/session"
	"context"
)

func CreateStreamResponse(ctx context.Context, s *session.Session, input string) (<-chan string, <-chan apitypes.Response, <-chan error, error) {
	return nil, nil, nil, nil
}
