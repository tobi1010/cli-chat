package providers

type providerFn func(string) Provider

var providersMap = map[string]providerFn{
	"openai":    OpenaiFn,
	"anthropic": AnthropicFn,
}

func Get(name string, model string) (Provider, bool) {
	prvFn, ok := providersMap[name]
	if !ok {
		return Provider{}, false
	}
	return prvFn(model), true
}
