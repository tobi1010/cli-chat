package index

import "time"

type DB struct {
	Chats []ChatMeta `json:"chats"`
}

func NewDB() *DB {
	return &DB{
		Chats: []ChatMeta{},
	}
}

type ChatMeta struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
