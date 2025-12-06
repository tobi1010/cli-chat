package fileatomic

import (
	"fmt"
	"os"
	"path/filepath"
)

func Write(path string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(path)
	err := os.MkdirAll(dir, 0o700)
	if err != nil {
		return fmt.Errorf("creating dir %s: %w", dir, err)
	}
	tmp := path + ".tmp"
	fd, err := os.OpenFile(tmp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
	if err != nil {
		return fmt.Errorf("opening file %s: %w", tmp, err)
	}
	_, err = fd.Write(data)
	if err != nil {
		_ = fd.Close()
		return fmt.Errorf("writing data to %s: %w", tmp, err)
	}
	err = fd.Sync()
	if err != nil {
		_ = fd.Close()
		return fmt.Errorf("syncing file: %w", err)
	}
	err = fd.Close()
	if err != nil {
		return fmt.Errorf("closing file %s: %w", tmp, err)
	}
	err = os.Rename(tmp, path)
	if err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("renaming file %s: %w", path, err)
	}
	return nil
}
