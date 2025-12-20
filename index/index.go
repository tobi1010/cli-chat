package index

import (
	"time"
)

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

func (db *DB) Touch(chatId string, now time.Time) {
	if meta, ok := db.Find(chatId); ok {
		meta.UpdatedAt = now
		return
	}
	db.Chats = append(db.Chats, ChatMeta{ID: chatId, CreatedAt: now, UpdatedAt: now})
}

func (db *DB) Find(chatId string) (*ChatMeta, bool) {
	for i := range db.Chats {
		if db.Chats[i].ID == chatId {
			return &db.Chats[i], true
		}
	}
	return nil, false
}

func (db *DB) GetLastChatId() string {
	if db == nil || len(db.Chats) == 0 {
		return ""
	}

	last := db.Chats[0]
	for i := 1; i < len(db.Chats); i++ {
		c := db.Chats[i]

		if c.UpdatedAt.After(last.UpdatedAt) {
			last = c
			continue
		}
		if c.UpdatedAt.Equal(last.UpdatedAt) && c.ID > last.ID {
			last = c
		}
	}

	return last.ID
}
