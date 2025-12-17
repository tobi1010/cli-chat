package fileatomic

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"
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

	if err := Write(fp, data, perm); err != nil {
		t.Fatalf("writing file: %v", err)
	}

	info, err := os.Stat(fp)
	if err != nil {
		t.Fatalf("stat: %v", err)
	}

	if tmpFilePerm := info.Mode().Perm(); tmpFilePerm != perm {
		t.Fatalf("permissions mismatch. want: %o, got: %o", perm, tmpFilePerm)
	}

	if got := readFile(t, fp); !bytes.Equal(data, got) {
		t.Fatalf("data mismatch: want %q, got %q", data, got)
	}

	entries, err := os.ReadDir(tmpDir)
	if err != nil {
		t.Fatalf("reading temp dir: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("unexpected number of files in temp dir: %v", entries)
	}

}

func TestAtomicWrite_Concurrent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "file.txt")

	const writers = 50
	var wg sync.WaitGroup
	expected := make(map[string]struct{}, writers)
	var mu sync.Mutex

	for i := range writers {
		wg.Add(1)

		go func(j int) {
			defer wg.Done()

			data := fmt.Appendf(nil, "writer-%d", j)

			if err := Write(path, data, 0o644); err != nil {
				t.Errorf("write failed: %v", err)
				return
			}

			mu.Lock()
			expected[string(data)] = struct{}{}
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	final, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("final read failed: %v", err)
	}

	if _, ok := expected[string(final)]; !ok {
		t.Fatalf("final content not from any writer: %q", final)
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Fatalf("temp files leaked: %v", entries)
	}
}

func TestOverwriteExistingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.txt")
	originalData := []byte("original data")
	originalPerm := os.FileMode(0o600)
	if err := os.WriteFile(path, originalData, originalPerm); err != nil {
		t.Fatalf("writing original file: %v", err)
	}
	newData := []byte("new data")
	newPerm := os.FileMode(0o644)
	if err := Write(path, newData, newPerm); err != nil {
		t.Fatalf("writing file atomically: %v", err)
	}
	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading file: %v", err)
	}
	if !bytes.Equal(got, newData) {
		t.Fatalf("got: %q, wanted: %q", got, newData)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat: %v", err)
	}
	actualPerm := info.Mode().Perm()
	if actualPerm != newPerm {
		t.Fatalf("permissions mismatch. want: %o, got: %o", newPerm, actualPerm)
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 {
		t.Fatalf("temp files leaked: %v", entries)
	}

}

func TestWrite_CreatesParentDir(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "a", "b", "c", "file.txt")

	data := []byte("hello")
	perm := os.FileMode(0o644)

	if err := Write(path, data, perm); err != nil {
		t.Fatalf("Write: %v", err)
	}

	got, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile: %v", err)
	}
	if !bytes.Equal(got, data) {
		t.Fatalf("content mismatch: got %q, want %q", got, data)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat: %v", err)
	}
	if info.Mode().Perm() != perm {
		t.Fatalf("perm mismatch: got %o, want %o", info.Mode().Perm(), perm)
	}

	entries, err := os.ReadDir(filepath.Dir(path))
	if err != nil {
		t.Fatalf("ReadDir: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("temp files leaked: %v", entries)
	}
}

func TestWrite_FailsWhenParentIsFile(t *testing.T) {
	dir := t.TempDir()

	parent := filepath.Join(dir, "parent")
	if err := os.WriteFile(parent, []byte("not a dir"), 0o644); err != nil {
		t.Fatalf("seed parent file: %v", err)
	}

	path := filepath.Join(parent, "child.txt")
	err := Write(path, []byte("data"), 0o644)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
