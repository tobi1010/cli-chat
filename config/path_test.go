package config

import (
	"path/filepath"
	"testing"
)

func TestGetConfigPath_Uses_XDG(t *testing.T) {
	t.Run("uses XDG_CONFIG_HOME when set", func(t *testing.T) {
		tmp := t.TempDir()
		t.Setenv("XDG_CONFIG_HOME", tmp)
		got, err := GetSettingsPath()
		if err != nil {
			t.Fatalf("GetSettingsPath: %v", err)
		}

		want := filepath.Join(tmp, APP_DIR)
		if got != want {
			t.Fatalf("wanted %q but got %q", want, got)
		}
	})
}
