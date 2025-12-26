package providers

type Provider struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	BaseURL string `json:"baseurl"`
	Model   string `json:"model"`
}

func NewDefault() Provider {
	provider := OpenaiFn("")
	return provider
}

func OpenaiFn(model string) Provider {
	provider := Provider{
		Name:    "openai",
		Key:     "OPENAI_API_KEY",
		BaseURL: "https://api.openai.com/v1/",
		Model:   "gpt-5",
	}
	if model != "" {
		provider.Model = model
	}
	return provider
}

func AnthropicFn(model string) Provider {
	provider := Provider{
		Name:    "anthropic",
		Key:     "ANTHROPIC_API_KEY",
		BaseURL: "https://api.anthropic.com/v1/",
		Model:   "gpt-5",
	}
	if model != "" {
		provider.Model = model
	}
	return provider
}
