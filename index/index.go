package index

import (
	"time"
)

type DB struct {
	Chats []ChatMeta `json:"chats"`

	ChatsDir  string `json:"-"`
	IndexPath string `json:"-"`
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

func (db *DB) Touch(chatId string, now time.Time) {
	if meta, ok := db.find(chatId); ok {
		meta.UpdatedAt = now
		return
	}
	db.Chats = append(db.Chats, ChatMeta{ID: chatId, CreatedAt: now, UpdatedAt: now})
}

func (db *DB) find(chatId string) (*ChatMeta, bool) {
	for i := range db.Chats {
		if db.Chats[i].ID == chatId {
			return &db.Chats[i], true
		}
	}
	return nil, false
}
