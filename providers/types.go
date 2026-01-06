package providers

import (
	"cli-chat/internal/apitypes"
)

const Default = "openai"

type Provider struct {
	Name    string         `json:"name"`
	Key     string         `json:"key"`
	BaseURL string         `json:"baseurl"`
	Model   apitypes.Model `json:"model"`
}

type ProviderDef struct {
	Name         string
	EnvKey       string
	BaseURL      string
	DefaultModel string
}

var Registry = map[string]ProviderDef{
	"openai": {
		Name:         "openai",
		EnvKey:       "OPENAI_API_KEY",
		BaseURL:      "https://api.openai.com/v1/",
		DefaultModel: "gpt-5",
	},
	"anthropic": {
		Name:         "anthropic",
		EnvKey:       "ANTHROPIC_API_KEY",
		BaseURL:      "https://api.anthropic.com/v1/",
		DefaultModel: "claude-opus-4-5-20251101",
	},
}

func Def(name string) (ProviderDef, bool) {
	def, ok := Registry[name]
	return def, ok
}
func Names() []string {
	names := make([]string, 0, len(Registry))
	for k := range Registry {
		names = append(names, k)
	}
	return names
}
