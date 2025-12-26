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
