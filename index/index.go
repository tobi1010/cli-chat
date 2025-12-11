package index

import (
	"time"
)

type DB struct {
	Chats     []ChatMeta
	ChatsDir  string
	IndexPath string
}

func NewDB(indexPath string, chatsDir string) (*DB, error) {
	return &DB{
		Chats:     []ChatMeta{},
		ChatsDir:  chatsDir,
		IndexPath: indexPath,
	}, nil
}

type ChatMeta struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
