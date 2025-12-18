package session

import (
	"cli-chat/paths"
	"os"
	"testing"
)

func TestLoadOrCreate(t *testing.T) {
	tests := []struct {
		name string
		init func(t *testing.T)
	}{
		{name: "first run, no session, settings or db",
			init: func(t *testing.T) {
				t.Setenv("XDG_CONFIG_HOME", t.TempDir())
				t.Setenv("XDG_DATA_HOME", t.TempDir())
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.init(t)
			sessionPath, err := paths.SessionPath()
			if err != nil {
				t.Fatalf("resolving session path: %v", err)
			}
			s, err := LoadOrCreate(sessionPath)
			if err != nil {
				t.Fatalf("LoadOrCreate session: %v", err)
			}

			if s.Client == nil {
				t.Fatalf("expected non-nil client")
			}
			if s.Chat == nil {
				t.Fatalf("expected non-nil chat")
			}
			if s.DB == nil {
				t.Fatalf("expected non-nil db")
			}

			settingsPath, err := paths.SettingsPath()
			if err != nil {
				t.Fatalf("resolving settings path: %v", err)
			}
			if _, err := os.Stat(settingsPath); err != nil {
				t.Fatalf("expected settings file to exist at %s: %v", settingsPath, err)
			}

			if _, err := os.Stat(sessionPath); err == nil {
				t.Fatalf("expected session file to not exist at %s", sessionPath)
			} else if !os.IsNotExist(err) {
				t.Fatalf("stat %s: %v", sessionPath, err)
			}

			dbFilePath, err := paths.IndexPath()
			if err != nil {
				t.Fatalf("resolving index file path: %v", err)
			}
			if _, err := os.Stat(dbFilePath); err == nil {
				t.Fatalf("expected index file to not exist at %s", dbFilePath)
			} else if !os.IsNotExist(err) {
				t.Fatalf("stat %s: %v", dbFilePath, err)
			}

		})
	}
}
