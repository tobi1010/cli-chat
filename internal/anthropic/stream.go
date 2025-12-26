package anthropic

func CreateStreamResponse(ctx context.Context, s *session.Session, input string) (<-chan string, <-chan llm.Response, <-chan error, error)
