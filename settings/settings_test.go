package settings

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPrintSettings_MasksApiKey(t *testing.T) {
	s := NewDefaultSettings()
	s.ApiKey = "super-secret"

	origStdout := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w
	t.Cleanup(func() {
		os.Stdout = origStdout
		_ = r.Close()
	})

	done := make(chan string, 1)
	go func() {
		var buf bytes.Buffer
		_, _ = io.Copy(&buf, r)
		done <- buf.String()
	}()

	require.NoError(t, s.PrintSettings())

	require.NoError(t, w.Close())
	out := <-done

	require.Contains(t, out, "Settings:")
	require.Contains(t, out, "api_key")
	require.Contains(t, out, "[set]")
	require.NotContains(t, out, "super-secret")
}
