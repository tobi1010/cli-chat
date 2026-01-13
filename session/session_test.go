package session

import (
	"os"
	"path/filepath"
	"terminal-chat/paths"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUpdateChat(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv("HOME", t.TempDir())

	p, err := paths.ResolvePaths()
	require.NoError(t, err)
	s, err := Open(p)
	require.NoError(t, err)
	require.NotNil(t, s)
	require.NotNil(t, s.Chat)
	require.NotNil(t, s.DB)

	s.Chat.ID = "chat1"
	s.Chat.UpdatedAt = time.Unix(123, 0)

	err = s.UpdateChat()
	require.NoError(t, err)

	indexPath, err := paths.IndexPath()
	require.NoError(t, err)
	_, err = os.Stat(indexPath)
	require.NoError(t, err)

	chatsDir, err := paths.ChatsDir()
	require.NoError(t, err)

	chatPath := filepath.Join(chatsDir, s.Chat.ID+".json")
	_, err = os.Stat(chatPath)
	require.NoError(t, err)

	meta, ok := s.DB.Find(s.Chat.ID)
	require.True(t, ok)
	require.True(t, s.Chat.UpdatedAt.Equal(meta.UpdatedAt))
}
