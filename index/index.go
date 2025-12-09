package index

import (
	"time"
)

type DB struct {
	Chats []ChatMeta
}

type ChatMeta struct {
	ID           string    `json:"id"`
	Title        string    `json:"title"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Path         string    `json:"path"`
	MessageCount int       `json:"message_count"`
	TokenCount   int       `json:"token_count"`
}
