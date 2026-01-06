package chat

import (
	"cli-chat/fileatomic"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Load(chatsDir string, id string) (*Chat, error) {

	path := filepath.Join(chatsDir, id+".json")

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", path, err)
	}
	var chat Chat
	if err = json.Unmarshal(data, &chat); err != nil {
		return nil, fmt.Errorf("unmarshalling json: %w", err)
	}

	return &chat, nil
}

func (c *Chat) Save(chatsDir string) error {

	path := filepath.Join(chatsDir, c.ID+".json")

	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	if err = fileatomic.Write(path, data, 0o600); err != nil {
		return fmt.Errorf("writing chat atomically: %w", err)
	}

	return nil
}
