package providers

const Default = "openai"

type Provider struct {
	ID           string
	Name         string
	EnvKey       string
	BaseURL      string
	DefaultModel string
}

var Registry = map[string]Provider{
	"openai": {
		ID:           "openai",
		Name:         "OpenAi",
		EnvKey:       "OPENAI_API_KEY",
		BaseURL:      "https://api.openai.com/v1/",
		DefaultModel: "gpt-5",
	},
	"anthropic": {
		ID:           "anthropic",
		Name:         "Anthropic",
		EnvKey:       "ANTHROPIC_API_KEY",
		BaseURL:      "https://api.anthropic.com/v1/",
		DefaultModel: "claude-opus-4-5-20251101",
	},
}

func Get(id string) (Provider, bool) {
	def, ok := Registry[id]
	return def, ok
}

func DefaultNameAndModelID() (string, string) {
	return Registry[Default].Name, Registry[Default].DefaultModel
}

func Names() []string {
	names := make([]string, 0, len(Registry))
	for k := range Registry {
		names = append(names, k)
	}
	return names
}
