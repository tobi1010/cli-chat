package providers

type Provider struct {
	Name    string `json:"name"`
	Key     string `json:"key"`
	BaseURL string `json:"baseurl"`
	Model   string `json:"model"`
}

func NewDefault() Provider {
	provider := Gpt5()
	return provider
}

func Gpt5() Provider {
	return Provider{
		Name:    "openai",
		Key:     "OPENAI_API_KEY",
		BaseURL: "https://api.openai.com/v1/",
		Model:   "gpt-5",
	}
}
