package settings

import (
	"cli-chat/paths"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadOrCreate_Missing_CreatesDefaultsAndWrites(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv("HOME", t.TempDir())

	settingsPath, err := paths.SettingsPath()
	require.NoError(t, err)

	_, err = os.Stat(settingsPath)
	require.Error(t, err) // should not exist yet

	got, err := LoadOrCreate()
	require.NoError(t, err)

	want := NewDefaultSettings()
	require.Equal(t, want, got)

	_, err = os.Stat(settingsPath)
	require.NoError(t, err)
}

func TestLoadOrCreate_Existing_Loads(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv("HOME", t.TempDir())

	settingsPath, err := paths.SettingsPath()
	require.NoError(t, err)

	customSettings := NewDefaultSettings()
	customSettings.ApiKey = "secret"

	err = os.MkdirAll(filepath.Dir(settingsPath), 0o700)
	require.NoError(t, err)
	err = customSettings.Save(settingsPath)
	require.NoError(t, err)

	got, err := LoadOrCreate()
	require.NoError(t, err)

	require.Equal(t, customSettings, got)
}

func TestLoadOrCreate_EmptyFile_WritesDefaults(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv("HOME", t.TempDir())

	p, err := paths.SettingsPath()
	require.NoError(t, err)

	err = os.MkdirAll(filepath.Dir(p), 0o700)
	require.NoError(t, err)

	err = os.WriteFile(p, []byte{}, 0o600)
	require.NoError(t, err)

	got, err := LoadOrCreate()
	require.NoError(t, err)

	want := NewDefaultSettings()
	require.Equal(t, want, got)
}
