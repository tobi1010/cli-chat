package settings

import (
	"encoding/json"
	"fmt"
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

	keys := make([]string, 0, len(settingsMap))
	// determine key length for print formatting
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
	// iterate over sorted keys
	for _, key := range keys {
		val := settingsMap[key]
		if key == "api_key" && s.ApiKey != "" {
			//mask api-key
			val = "[set]"
		}
		fmt.Printf("  %-*s %v\n", pad, key, val)
	}
	return nil
}
