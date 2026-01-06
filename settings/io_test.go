package settings

import (
	"cli-chat/paths"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaveAndLoadSettings_Roundtrip(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv("HOME", t.TempDir())

	s := NewDefaultSettings()
	settingsPath, err := paths.SettingsPath()
	require.NoError(t, err)

	err = os.MkdirAll(filepath.Dir(settingsPath), 0o700)
	require.NoError(t, err)

	err = s.Save(settingsPath)
	require.NoError(t, err)

	got, err := Load(settingsPath)
	require.NoError(t, err)

	require.Equal(t, s, got)
}

func TestLoadSettings_Missing_ReturnsDefaults(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv("HOME", t.TempDir())

	settingsPath, err := paths.SettingsPath()
	require.NoError(t, err)

	got, err := Load(settingsPath)
	require.NoError(t, err)

	want := NewDefaultSettings()
	require.Equal(t, want, got)
}

func TestLoadSettings_InvalidJSON_ReturnsError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "settings.json")
	err := os.WriteFile(path, []byte("{not-json"), 0o600)
	require.NoError(t, err)

	_, err = Load(path)
	require.Error(t, err)
}
