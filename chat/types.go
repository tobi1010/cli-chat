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
	Provider     string    `json:"provider"`
	Model        string    `json:"model"`
	Conversation []Message `json:"conversation"`
}

func New(providerName string, modelName string) *Chat {
	return &Chat{
		ID:           uuid.NewString(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		Provider:     providerName,
		Model:        modelName,
		Conversation: make([]Message, 0, 20),
	}
}
