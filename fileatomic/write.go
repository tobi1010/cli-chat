package fileatomic

import (
	"fmt"
	"os"
	"path/filepath"
)

func Write(path string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("creating dir %s: %w", dir, err)
	}
	base := filepath.Base(path)
	f, err := os.CreateTemp(dir, base+".tmp-*")
	if err != nil {
		return fmt.Errorf("creating temp file in %s: %w", dir, err)
	}
	tmp := f.Name()
	defer func() {
		_ = f.Close()
		_ = os.Remove(tmp)
	}()

	if err = f.Chmod(perm); err != nil {
		return fmt.Errorf("chmod temp file %s: %w", tmp, err)
	}

	if _, err = f.Write(data); err != nil {
		return fmt.Errorf("writing data to %s: %w", tmp, err)
	}
	if err = f.Sync(); err != nil {
		return fmt.Errorf("syncing file: %w", err)
	}
	if err = f.Close(); err != nil {
		return fmt.Errorf("closing file %s: %w", tmp, err)
	}
	if err = os.Rename(tmp, path); err != nil {
		return fmt.Errorf("renaming file %s: %w", path, err)
	}

	if d, err := os.Open(dir); err == nil {
		_ = d.Sync()
		_ = d.Close()
	}

	return nil
}
