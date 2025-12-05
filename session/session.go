package session

import (
	"cli-chat/chat"
	"cli-chat/index"
	"fmt"
	"path/filepath"
)

type Session struct {
	Provider string
	Model    string
	ChatsDir string
	DbPath   string
	Chat     *chat.Chat
}

func New(chatsDir string, dbPath string) *Session {
	return &Session{
		ChatsDir: chatsDir,
		DbPath:   dbPath,
	}

}
func (s *Session) LoadOrCreate() error {
	db, err := index.Load(s.ChatsDir)
	if err != nil {
		return fmt.Errorf("loading index file: %w", err)
	}
	chat, err := chat.ReadChat(filepath.Join(s.ChatsDir + db.LastChatID + ".json"))
	s.Provider = chat.Provider
	s.Model = chat.Model
	return nil
}

func (s *Session) SaveAtomic() error {
	return nil
}
