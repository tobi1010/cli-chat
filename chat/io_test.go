package chat

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestAddMessage_EmptyChat(t *testing.T) {
	c := Chat{
		ID:           "1",
		CreatedAt:    time.Unix(100, 0),
		UpdatedAt:    time.Unix(110, 0),
		Conversation: []Message{},
	}

	message1 := Message{Role: "user", Content: "user msg 1"}
	message2 := Message{Role: "assistant", Content: "halucination 1"}
	want := Chat{
		ID:           "1",
		CreatedAt:    time.Unix(100, 0),
		UpdatedAt:    time.Unix(110, 0),
		Conversation: []Message{message1, message2},
	}
	before := time.Now()
	c.AddMessage(message1.Role, message1.Content)
	after := time.Now()

	require.False(t, c.UpdatedAt.Before(before), "UpdatedAt < before")
	require.False(t, c.UpdatedAt.After(after), "UpdatedAt > after")
	require.False(t, c.UpdatedAt.Before(c.CreatedAt), "UpdatedAt < CreatedAt")
	require.Len(t, c.Conversation, 1)
	require.Equal(t, message1, c.Conversation[0])

	prev := c.UpdatedAt
	before = time.Now()
	c.AddMessage(message2.Role, message2.Content)
	after = time.Now()
	require.False(t, c.UpdatedAt.Before(before), "UpdatedAt < before")
	require.False(t, c.UpdatedAt.Before(prev), "previous < before")
	require.False(t, c.UpdatedAt.After(after), "UpdatedAt > after")
	require.False(t, c.UpdatedAt.Before(c.CreatedAt), "UpdatedAt < CreatedAt")

	require.Len(t, c.Conversation, 2)
	require.Equal(t, message2, c.Conversation[1])
	require.Equal(t, want.Conversation, c.Conversation)

}

func TestChat_WriteAndLoad_RoundTrip(t *testing.T) {
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
	err := c.Write(dir)
	require.NoError(t, err)

	got, err := LoadChat(dir, "chat1")
	require.NoError(t, err)
	require.Equal(t, c, got)
}

func TestLoadChat_NotFound(t *testing.T) {
	dir := t.TempDir()
	_, err := LoadChat(dir, "missing")
	require.Error(t, err)
}
