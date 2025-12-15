package paths

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetConfigRoot(t *testing.T) {
	tests := []struct {
		name  string
		setup func(t *testing.T)
	}{
		{
			name: "XDG_CONFIG_HOME set",
			setup: func(t *testing.T) {
				tmp := t.TempDir()
				t.Setenv("XDG_CONFIG_HOME", tmp)
			},
		},
		{
			name: "XDG_CONFIG_HOME not set",
			setup: func(t *testing.T) {
				t.Setenv("XDG_CONFIG_HOME", "")
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)

			var want string
			switch tc.name {
			case "XDG_CONFIG_HOME set":
				{
					want = os.Getenv("XDG_CONFIG_HOME")
				}
			case "XDG_CONFIG_HOME not set":
				{
					home, err := os.UserHomeDir()
					if err != nil {
						t.Fatalf("resolving user home dir: %v", err)
					}
					want = filepath.Join(home, ".config")
				}
			}
			got, err := GetConfigRoot()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != want {
				t.Fatalf("\ngot: %q, \nwant: %q\n", got, want)
			}
		})
	}

}

func TestAppConfigDir(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	root, err := GetConfigRoot()
	if err != nil {
		t.Fatalf("no config root: %v", err)
	}
	want := filepath.Join(root, AppDirName)
	got, err := AppConfigDir()
	if err != nil {
		t.Fatalf("resolving app config dir: %v", err)
	}
	if got != want {
		t.Fatalf("\ngot: %q, \nwant: %q\n", got, want)
	}
}

func TestSettingsPath(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	cfgDir, err := AppConfigDir()
	if err != nil {
		t.Fatalf("resolving config dir: %v", err)
	}
	want := filepath.Join(cfgDir, SettingsFile)
	got, err := SettingsPath()
	if err != nil {
		t.Fatalf("resolving settings file path: %v", err)
	}
	if got != want {
		t.Fatalf("\ngot: %q, \nwant: %q\n", got, want)
	}
}

func TestAppDataDir(t *testing.T) {
	tests := []struct {
		name  string
		setup func(t *testing.T)
	}{
		{
			name: "XDG_DATA_HOME set",
			setup: func(t *testing.T) {
				tmp := t.TempDir()
				t.Setenv("XDG_DATA_HOME", tmp)
			},
		},
		{
			name: "XDG_DATA_HOME not set",
			setup: func(t *testing.T) {
				t.Setenv("XDG_DATA_HOME", "")
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(t)

			var want string
			switch tc.name {
			case "XDG_DATA_HOME set":
				{
					xdg := os.Getenv("XDG_DATA_HOME")
					want = filepath.Join(xdg, AppDirName)
				}
			case "XDG_DATA_HOME not set":
				{
					home, err := os.UserHomeDir()
					if err != nil {
						t.Fatalf("resolving user home dir: %v", err)
					}
					want = filepath.Join(home, ".local", "share", AppDirName)
				}
			}
			got, err := AppDataDir()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != want {
				t.Fatalf("\ngot: %q, \nwant: %q\n", got, want)
			}
		})
	}

}

func TestSessionPath(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	xdg := os.Getenv("XDG_DATA_HOME")
	want := filepath.Join(xdg, AppDirName, SessionFile)
	got, err := SessionPath()
	if err != nil {
		t.Fatalf("resolving session file path: %v", err)
	}
	if got != want {
		t.Fatalf("\ngot: %q, \nwant: %q\n", got, want)
	}
}

func TestIndexPath(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	xdg := os.Getenv("XDG_DATA_HOME")
	want := filepath.Join(xdg, AppDirName, IndexFile)
	got, err := IndexPath()
	if err != nil {
		t.Fatalf("resolving index file path: %v", err)
	}
	if got != want {
		t.Fatalf("\ngot: %q, \nwant: %q\n", got, want)
	}
}

func TestChatsDir(t *testing.T) {
	t.Setenv("XDG_DATA_HOME", t.TempDir())
	xdg := os.Getenv("XDG_DATA_HOME")
	want := filepath.Join(xdg, AppDirName, ChatsDirName)
	got, err := ChatsDir()
	if err != nil {
		t.Fatalf("resolving chats dir: %v", err)
	}
	if got != want {
		t.Fatalf("\ngot: %q, \nwant: %q\n", got, want)
	}
}
