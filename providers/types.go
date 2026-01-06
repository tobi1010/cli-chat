package providers

import (
	"cli-chat/cache"
	"cli-chat/internal/apitypes"
	"fmt"
)

type Provider struct {
	Name    string         `json:"name"`
	Key     string         `json:"key"`
	BaseURL string         `json:"baseurl"`
	Model   apitypes.Model `json:"model"`
}

type Def struct {
	Name         string
	EnvKey       string
	BaseURL      string
	DefaultModel string
}

var Registry = map[string]Def{
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

func NewDefault(cache *cache.Cache) (Provider, error) {
	return New(cache, "openai", "")
}

func OpenaiFn(cache *cache.Cache, model string) (Provider, error) {
	return New(cache, "openai", model)
}

func AnthropicFn(cache *cache.Cache, model string) (Provider, error) {
	return New(cache, "anthropic", model)
}

func New(cache *cache.Cache, name string, model string) (Provider, error) {
	def, ok := Registry[name]
	if !ok {
		return Provider{}, fmt.Errorf("unknown provider %q", name)
	}

	p := Provider{
		Name:    def.Name,
		Key:     def.EnvKey,
		BaseURL: def.BaseURL,
	}

	requested := model
	if model == "" {
		model = def.DefaultModel
	}

	m, ok := cache.Get(def.Name, model)
	if ok {
		p.Model = m
		return p, nil
	}

	m, ok = cache.Get(def.Name, def.DefaultModel)
	if !ok {
		return Provider{}, fmt.Errorf(
			"unknown %s model %q and default %q not in cache",
			def.Name,
			requested,
			def.DefaultModel,
		)
	}

	p.Model = m
	return p, fmt.Errorf(
		"unknown %s model %q (using %q)",
		def.Name,
		requested,
		def.DefaultModel,
	)
}
