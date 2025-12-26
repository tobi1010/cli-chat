package settings

import (
	"cli-chat/fileatomic"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func LoadSettings(settingsPath string) (Settings, error) {
	data, err := os.ReadFile(settingsPath)
	var settings Settings
	if err != nil {
		if os.IsNotExist(err) {
			settings = NewDefaultSettings()
			return settings, nil
		}
		return Settings{}, fmt.Errorf("reading config file: %w", err)
	}
	if err = json.Unmarshal(data, &settings); err != nil {
		return Settings{}, fmt.Errorf("unmarshalling json: %w", err)
	}
	return settings, nil
}

func (s *Settings) Save(settingsPath string) error {
	dir := filepath.Dir(settingsPath)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return fmt.Errorf("creating settings dir: %w", err)
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling config: %w", err)
	}
	if err = fileatomic.Write(settingsPath, data, 0o600); err != nil {
		return fmt.Errorf("writing settings: %w", err)
	}
	return nil
}
