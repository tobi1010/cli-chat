package fileatomic

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func readFile(t *testing.T, fp string) []byte {
	t.Helper()
	data, err := os.ReadFile(fp)
	if err != nil {
		t.Fatalf("reading file %s: %v", fp, err)
	}
	return data
}

func TestWrite(t *testing.T) {
	tmpDir := t.TempDir()
	fp := filepath.Join(tmpDir, "testfile.txt")
	data := []byte("Hello amazing world of atomic file writing!")
	perm := os.FileMode(0o644)
	err := Write(fp, data, perm)
	if err != nil {
		t.Fatalf("writing file: %v", err)
	}
	info, err := os.Stat(fp)
	if os.IsNotExist(err) {
		t.Fatalf("file does not exist %s: %v", fp, err)
	}
	if err != nil {
		t.Fatalf("stat: %v", err)
	}
	tmpFilePerm := info.Mode().Perm()
	if tmpFilePerm != perm {
		t.Fatalf("permissions mismatch. want: %d, got: %d", perm, tmpFilePerm)
	}

	got := readFile(t, fp)
	if !bytes.Equal(data, got) {
		t.Fatalf("data mismatch: want %q, got %q", data, got)
	}

}
