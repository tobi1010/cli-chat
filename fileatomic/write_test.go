package fileatomic

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func readFile(t *testing.T, fp string) []byte {
	t.Helper()
	data, err := os.ReadFile(fp)
	require.NoError(t, err)
	return data
}

func TestWrite(t *testing.T) {
	tmpDir := t.TempDir()
	fp := filepath.Join(tmpDir, "file.txt")
	data := []byte("Hello amazing world of atomic file writing!")
	perm := os.FileMode(0o644)

	err := Write(fp, data, perm)
	require.NoError(t, err)

	info, err := os.Stat(fp)
	require.NoError(t, err)

	tmpFilePerm := info.Mode().Perm()
	require.Equal(t, perm, tmpFilePerm)

	got := readFile(t, fp)
	require.Equal(t, data, got)

	entries, err := os.ReadDir(tmpDir)
	require.NoError(t, err)
	require.Len(t, entries, 1)
	require.Equal(t, "file.txt", entries[0].Name())
}

func TestAtomicWrite_Concurrent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "file.txt")

	const writers = 50
	var wg sync.WaitGroup
	expected := make(map[string]struct{}, writers)
	var mu sync.Mutex

	errCh := make(chan error, writers)

	for i := range writers {
		wg.Add(1)

		go func(j int) {
			defer wg.Done()

			data := fmt.Appendf(nil, "writer-%d", j)

			errCh <- Write(path, data, 0o644)

			mu.Lock()
			expected[string(data)] = struct{}{}
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		require.NoError(t, err)
	}

	final, err := os.ReadFile(path)
	require.NoError(t, err)

	_, ok := expected[string(final)]
	require.True(t, ok)

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, entries, 1)
	require.Equal(t, "file.txt", entries[0].Name())
}

func TestOverwriteExistingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "file.txt")
	originalData := []byte("original data")
	originalPerm := os.FileMode(0o600)

	err := os.WriteFile(path, originalData, originalPerm)
	require.NoError(t, err)

	wantData := []byte("new data")
	wantPerm := os.FileMode(0o644)

	err = Write(path, wantData, wantPerm)
	require.NoError(t, err)

	gotData, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, wantData, gotData)

	info, err := os.Stat(path)
	require.NoError(t, err)

	gotPerm := info.Mode().Perm()
	require.Equal(t, wantPerm, gotPerm)

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	require.Len(t, entries, 1)
	require.Equal(t, "file.txt", entries[0].Name())
}

func TestWrite_CreatesParentDir(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "a", "b", "c", "file.txt")

	wantData := []byte("hello")
	wantPerm := os.FileMode(0o644)

	err := Write(path, wantData, wantPerm)
	require.NoError(t, err)

	gotData, err := os.ReadFile(path)
	require.NoError(t, err)

	require.Equal(t, wantData, gotData)

	info, err := os.Stat(path)
	require.NoError(t, err)

	gotPerm := info.Mode().Perm()
	require.Equal(t, wantPerm, gotPerm)

	entries, err := os.ReadDir(filepath.Dir(path))
	require.NoError(t, err)
	require.Len(t, entries, 1)
	require.Equal(t, "file.txt", entries[0].Name())
}

func TestWrite_FailsWhenParentIsFile(t *testing.T) {
	dir := t.TempDir()

	parent := filepath.Join(dir, "parent")
	err := os.WriteFile(parent, []byte("not a dir"), 0o644)
	require.NoError(t, err)

	path := filepath.Join(parent, "child.txt")
	err = Write(path, []byte("data"), 0o644)
	require.Error(t, err)
}
