package index

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func Load(path string) (*DB, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &DB{
			LastChatID: "",
			Chats:      make([]ChatMeta, 0, 10),
		}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("reading file %s: %w", path, err)
	}
	var db DB
	err = json.Unmarshal(data, &db)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling json: %w", err)
	}
	return &db, nil
}
func SaveAtomic(path string, db *DB) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0o700)
	if err != nil {
		return fmt.Errorf("creating dir %s: %w", dir, err)
	}
	tmp := path + ".tmp"
	tmpFile, err := os.OpenFile(tmp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o600)
	if err != nil {
		return fmt.Errorf("opening file %s: %w", tmp, err)
	}
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		tmpFile.Close()
		return fmt.Errorf("marshalling json: %w", err)
	}
	_, err = tmpFile.Write(data)
	if err != nil {
		tmpFile.Close()
		return fmt.Errorf("writing file %s: %w", tmpFile.Name(), err)
	}
	err = tmpFile.Sync()
	if err != nil {
		tmpFile.Close()
		return fmt.Errorf("syncing file: %w", err)
	}
	err = tmpFile.Close()
	if err != nil {
		return fmt.Errorf("closing file: %w", err)
	}

	err = os.Rename(tmp, path)
	if err != nil {
		return fmt.Errorf("renaming file %s: %w", path, err)
	}
	return nil
}
