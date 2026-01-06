package env

import (
	"fmt"
	"os"
	"strings"
)

func ResolveAPIKey(envVarName string) (string, error) {
	key := os.Getenv(envVarName)
	if key == "" {
		return "", fmt.Errorf("resolving API key. Make sure your API key is exported as <provider_name>_API_KEY!")
	}
	return key, nil
}

func FindApiKeyCandidates() []string {
	globalvars := os.Environ()
	candidates := []string{}
	for _, kv := range globalvars {
		parts := strings.SplitN(kv, "=", 2)
		key := parts[0]
		if strings.Contains(key, "API_KEY") {
			candidates = append(candidates, key)
		}
	}
	return candidates
}
