package auth

import (
	"fmt"
	"os"
)

func ResolveAPIKey(envVarName string) (string, error) {
	key := os.Getenv(envVarName)
	if key == "" {
		return "", fmt.Errorf("resolving API key. Make sure your API key is exported as <provider_name>_API_KEY!")
	}
	return key, nil
}
