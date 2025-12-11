package settings

import (
	"cli-chat/fileatomic"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func (s *Settings) PrintSettings() error {
	b, err := json.Marshal(s)
	if err != nil {
		return fmt.Errorf("marshal settings: %w", err)
	}
	var settingsMap map[string]any
	if err := json.Unmarshal(b, &settingsMap); err != nil {
		return fmt.Errorf("unmarshal to map: %w", err)
	}

	// Stable order
	keys := make([]string, 0, len(settingsMap))
	maxKeyLen := 0
	for k := range settingsMap {
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)
	pad := maxKeyLen + 3

	fmt.Println("Settings:")
	for _, k := range keys {
		v := settingsMap[k]
		// redact obvious secrets
		switch k {
		case "api_key", "token", "password", "secret":
			if vs, ok := v.(string); ok && vs != "" {
				fmt.Printf("  %-*s [set]\n", pad, k)
			} else {
				fmt.Printf("  %-*s [empty]\n", pad, k)
			}
			continue
		}
		fmt.Printf("  %-*s %v\n", pad, k, v)
	}
	return nil
}

func ReadSettings(settingsPath string) (*Settings, error) {
	data, err := os.ReadFile(settingsPath)
	var settings Settings
	if err != nil {
		if os.IsNotExist(err) {
			settings = NewDefaultSettings()
			return &settings, nil
		}
		return nil, fmt.Errorf("reading config file: %w", err)
	}
	err = json.Unmarshal(data, &settings)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling json: %w", err)
	}
	return &settings, nil
}

func (s *Settings) Save(settingsPath string) error {
	dir := filepath.Dir(settingsPath)
	err := os.MkdirAll(dir, 0o700)
	if err != nil {
		return fmt.Errorf("creating settings dir: %w", err)
	}
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("marshalling config: %w", err)
	}
	err = fileatomic.Write(settingsPath, data, 0o600)
	if err != nil {
		return fmt.Errorf("writing settings: %w", err)
	}
	return nil
}
