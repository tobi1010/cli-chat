package chat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestChat_SaveAndLoad_RoundTrip(t *testing.T) {
	dir := t.TempDir()

	c := &Chat{
		ID:        "chat1",
		CreatedAt: time.Unix(100, 0),
		UpdatedAt: time.Unix(110, 0),
		Conversation: []Message{
			{Role: "user", Content: "hi"},
			{Role: "assistant", Content: "yo"},
		},
	}
	err := c.Save(dir)
	require.NoError(t, err)

	got, err := Load(dir, "chat1")
	require.NoError(t, err)
	require.Equal(t, c, got)
}

func TestLoadChat_NotFound(t *testing.T) {
	dir := t.TempDir()
	_, err := Load(dir, "missing")
	require.Error(t, err)
}
