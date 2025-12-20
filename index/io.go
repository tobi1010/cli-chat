package index

import (
	"cli-chat/fileatomic"
	"encoding/json"
	"fmt"
	"os"
)

func Load(path string) (*DB, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &DB{
			Chats: make([]ChatMeta, 0, 10),
		}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", path, err)
	}
	var db DB
	if err = json.Unmarshal(data, &db); err != nil {
		return nil, fmt.Errorf("unmarshalling json: %w", err)
	}
	return &db, nil
}

func (db *DB) Save(path string) error {
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling json: %w", err)
	}
	if err = fileatomic.Write(path, data, 0o600); err != nil {
		return fmt.Errorf("writing atomically file %s: %w", path, err)
	}
	return nil
}
