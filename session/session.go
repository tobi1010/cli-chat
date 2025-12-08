package session

import (
	"cli-chat/chat"
	"cli-chat/index"
	"cli-chat/paths"
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

func Create() (*Session, error) {
	chatsDir, err := paths.ChatsDir()
	if err != nil {
		return nil, fmt.Errorf("resolving chats dir: %w", err)
	}

	idxFile, err := paths.IndexPath()
	if err != nil {
		return nil, fmt.Errorf("resolving inex file path: %w", err)
	}

	s := Session{
		ChatsDir: chatsDir,
		DbPath:   idxFile,
	}

	db, err := index.Load(s.DbPath)
	if err != nil {
		return nil, fmt.Errorf("reading index file: %w", err)
	}

	lastChatId := db.LastChatID
	lastChatPath := filepath.Join(s.ChatsDir, lastChatId+".json")
	chat, err := db.Load(lastChatPath)
	if err != nil {
		return nil, fmt.Errorf("loading chat %s: %w", &lastChatPath, err)
	}
	s.Chat = chat

	s.Provider = s.Chat.Provider
	s.Model = s.Chat.Model
	return &s, nil
}
