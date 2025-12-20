package index

import (
	"cli-chat/paths"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDBWriteLoad_Roundtrip(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv("HOME", t.TempDir())

	db := NewDB()

	indexPath, err := paths.IndexPath()
	require.NoError(t, err)

	db.Touch("1", time.Unix(100, 0))

	require.NoError(t, db.Save(indexPath))

	got, err := Load(indexPath)
	require.NoError(t, err)

	require.Equal(t, db, got)
}

func TestDBTouchFind(t *testing.T) {
	db := NewDB()

	db.Touch("1", time.Unix(100, 0))

	meta, ok := db.Find("1")
	require.True(t, ok)
	require.True(t, meta.CreatedAt.Equal(time.Unix(100, 0)))
	require.True(t, meta.UpdatedAt.Equal(time.Unix(100, 0)))
	require.Len(t, db.Chats, 1)

	db.Touch("1", time.Unix(200, 0))

	meta, ok = db.Find("1")
	require.True(t, ok)
	require.True(t, meta.CreatedAt.Equal(time.Unix(100, 0)))
	require.True(t, meta.UpdatedAt.Equal(time.Unix(200, 0)))
	require.Len(t, db.Chats, 1)

	db.Touch("2", time.Unix(300, 0))

	meta2, ok := db.Find("2")
	require.True(t, ok)
	require.True(t, meta2.CreatedAt.Equal(time.Unix(300, 0)))
	require.True(t, meta2.UpdatedAt.Equal(time.Unix(300, 0)))
	require.Len(t, db.Chats, 2)
}

func TestGetLastChatId(t *testing.T) {
	var db *DB
	require.Equal(t, "", db.GetLastChatId())

	db = NewDB()

	db.Touch("1", time.Unix(100, 0))
	db.Touch("2", time.Unix(200, 0))

	lastId := db.GetLastChatId()
	require.Equal(t, "2", lastId)

	db.Touch("3", time.Unix(300, 0))

	lastId = db.GetLastChatId()
	require.Equal(t, "3", lastId)

	db.Touch("4", time.Unix(300, 0))

	lastId = db.GetLastChatId()
	require.Equal(t, "4", lastId)
}
