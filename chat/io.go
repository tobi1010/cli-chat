package chat

import (
	"cli-chat/fileatomic"
	"cli-chat/paths"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func (c *Chat) AddMessage(role string, content string) {
	msg := Message{
		Role:    role,
		Content: content,
	}
	c.Conversation = append(c.Conversation, msg)
	c.UpdatedAt = time.Now()
}

func (c *Chat) Write() error {
	dir, err := paths.ChatsDir()
	if err != nil {
		return fmt.Errorf("resolving chats dir: %w", err)
	}

	path := filepath.Join(dir, c.ID+".json")

	err = os.MkdirAll(dir, 0o700)
	if err != nil {
		return fmt.Errorf("creating chat dir: %w", err)
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	err = fileatomic.Write(path, data, 0o600)
	if err != nil {
		return fmt.Errorf("writing chat atomically: %w", err)
	}

	return nil
}

func ReadChat(id string) (*Chat, error) {
	dir, err := paths.ChatsDir()
	if err != nil {
		return nil, fmt.Errorf("resolving chats dir: %w", err)
	}

	path := filepath.Join(dir, id+".json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file: %w", err)
	}
	var chat Chat
	err = json.Unmarshal(data, &chat)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling json: %w", err)
	}

	return &chat, nil
}
