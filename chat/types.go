package chat

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Chat struct {
	ID           string    `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Conversation []Message `json:"conversation"`
}

func New() *Chat {
	return &Chat{
		ID:           uuid.NewString(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Conversation: make([]Message, 0, 20),
	}
}
