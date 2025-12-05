package chat

import (
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

	f := filepath.Join(dir, c.ID+".json")

	err = os.MkdirAll(dir, 0o700)
	if err != nil {
		return fmt.Errorf("creating chat dir: %w", err)
	}
	tmpFile := f + ".tmp"
	fd, err := os.OpenFile(tmpFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return fmt.Errorf("opening temp file: %w", err)
	}

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		_ = fd.Close()
		return fmt.Errorf("marshalling json: %w", err)
	}

	_, err = fd.Write(data)
	if err != nil {
		_ = fd.Close()
		return fmt.Errorf("writing to tmp file: %w", err)
	}
	err = fd.Sync()
	if err != nil {
		_ = fd.Close()
		return fmt.Errorf("syncing to disk: %w", err)
	}
	err = fd.Close()
	if err != nil {
		return fmt.Errorf("closing file: %w", err)
	}

	err = os.Rename(tmpFile, f)
	if err != nil {
		_ = os.Remove(tmpFile)
		return fmt.Errorf("renaming file: %w", err)
	}
	return nil
}

func ReadChat(id string) (*Chat, error) {
	dir, err := paths.ChatsDir()
	if err != nil {
		return nil, fmt.Errorf("resolving chats dir: %w", err)
	}

	f := filepath.Join(dir, id+".json")

	data, err := os.ReadFile(f)
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
