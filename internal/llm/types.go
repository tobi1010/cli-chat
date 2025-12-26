package llm

import (
	"cli-chat/session"
	"context"
)

type Model struct {
	ID      string
	Created int
	OwnedBy string
}

type Response struct {
	Model  string
	Output []struct {
		Type    string `json:"type"`
		ID      string `json:"id"`
		Status  string `json:"status"`
		Role    string `json:"role"`
		Content []struct {
			Type        string `json:"type"`
			Text        string `json:"text"`
			Annotations []any  `json:"annotations"`
		} `json:"content"`
	} `json:"output"`
}

type Backend interface {
	GetModels(ctx context.Context, s *session.Session) ([]Model, error)
	CreateStreamResponse(
		ctx context.Context,
		s *session.Session,
		input string,
	) (<-chan string, <-chan Response, <-chan error, error)
}
