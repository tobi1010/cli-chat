package index

import (
	"terminal-chat/chat"
	"terminal-chat/paths"
	"fmt"
	"time"
)

func (db *DB) Touch(chatId string, now time.Time) {
	if meta, ok := db.Find(chatId); ok {
		meta.UpdatedAt = now
		return
	}
	db.Chats = append(db.Chats, ChatMeta{ID: chatId, CreatedAt: now, UpdatedAt: now})
}
func (db *DB) GetByID(chatId string) (*chat.Chat, error) {
	chatsDir, err := paths.ChatsDir()
	if err != nil {
		return nil, fmt.Errorf("resolving chats dir: %w", err)
	}
	loadedChat, err := chat.Load(chatsDir, chatId)
	if err != nil {
		return nil, fmt.Errorf("loading chat: %w", err)
	}
	return loadedChat, nil

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
