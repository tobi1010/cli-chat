package providers

const Default = "openai"

type Provider struct {
	Name         string
	EnvKey       string
	BaseURL      string
	DefaultModel string
}

var Registry = map[string]Provider{
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

func Get(name string) (Provider, bool) {
	def, ok := Registry[name]
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
