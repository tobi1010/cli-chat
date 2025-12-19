package paths

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetConfigRoot(t *testing.T) {
	tests := []struct {
		name  string
		setup func(t *testing.T)
		want  func(t *testing.T) string
	}{
		{
			name: "XDG_CONFIG_HOME set",
			setup: func(t *testing.T) {
				tmp := t.TempDir()
				t.Setenv("XDG_CONFIG_HOME", tmp)
			},
			want: func(t *testing.T) string {
				return os.Getenv("XDG_CONFIG_HOME")
			},
		},
		{
			name: "XDG_CONFIG_HOME not set",
			setup: func(t *testing.T) {
				t.Setenv("XDG_CONFIG_HOME", "")
			},
			want: func(t *testing.T) string {
				home, err := os.UserHomeDir()
				require.NoError(t, err)
				return filepath.Join(home, ".config")
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
			want := tc.want(t)

			got, err := GetConfigRoot()
			require.NoError(t, err)
			require.Equal(t, want, got)
		})
	}

}

func TestAppConfigDir(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	root, err := GetConfigRoot()
	require.NoError(t, err)
	want := filepath.Join(root, AppDirName)
	got, err := AppConfigDir()
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestSettingsPath(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	cfgDir, err := AppConfigDir()
	require.NoError(t, err)
	want := filepath.Join(cfgDir, SettingsFile)
	got, err := SettingsPath()
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestAppDataDir(t *testing.T) {
	tests := []struct {
		name  string
		setup func(t *testing.T)
		want  func(t *testing.T) string
	}{
		{
			name: "XDG_DATA_HOME set",
			setup: func(t *testing.T) {
				tmp := t.TempDir()
				t.Setenv("XDG_DATA_HOME", tmp)
			},
			want: func(t *testing.T) string {
				xdg := os.Getenv("XDG_DATA_HOME")
				return filepath.Join(xdg, AppDirName)
			},
		},
		{
			name: "XDG_DATA_HOME not set",
			setup: func(t *testing.T) {
				t.Setenv("XDG_DATA_HOME", "")
			},
			want: func(t *testing.T) string {
				home, err := os.UserHomeDir()
				require.NoError(t, err)
				return filepath.Join(home, ".local", "share", AppDirName)
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)
			want := tc.want(t)

			got, err := AppDataDir()
			require.NoError(t, err)
			require.Equal(t, want, got)
		})
	}

}

func TestSessionPath(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	xdg := os.Getenv("XDG_DATA_HOME")
	want := filepath.Join(xdg, AppDirName, SessionFile)
	got, err := SessionPath()
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestIndexPath(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	xdg := os.Getenv("XDG_DATA_HOME")
	want := filepath.Join(xdg, AppDirName, IndexFile)
	got, err := IndexPath()
	require.NoError(t, err)
	require.Equal(t, want, got)
}

func TestChatsDir(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	xdg := os.Getenv("XDG_DATA_HOME")
	want := filepath.Join(xdg, AppDirName, ChatsDirName)
	got, err := ChatsDir()
	require.NoError(t, err)
	require.Equal(t, want, got)
}
