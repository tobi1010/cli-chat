package anthropic

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"terminal-chat/internal/apitypes"
	"terminal-chat/internal/client"
	"terminal-chat/internal/env"
	"terminal-chat/providers"
)

func GetModels(ctx context.Context, client *client.Client, providerName string) ([]apitypes.Model, error) {
	provider, ok := providers.Get(providerName)
	if !ok {
		return nil, fmt.Errorf("unknown provider %q", providerName)
	}
	apiKey, err := env.ResolveAPIKey(provider.EnvKey)
	if err != nil {
		return nil, fmt.Errorf("resloving API key for %s: %w", provider.Name, err)
	}
	url := provider.BaseURL + "models"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("X-Api-Key", apiKey)

	res, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("doing request to %s: %w", url, err)
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("bad status: %s", res.Status)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body: %w", err)
	}
	var modelsRes AnthropicModelsResponse
	if err = json.Unmarshal(data, &modelsRes); err != nil {
		return nil, fmt.Errorf("unmarshalling json: %w", err)
	}
	models := modelsRes.Data
	apiModels := []apitypes.Model{}
	for _, m := range models {
		apiModel, err := toApiModel(m)
		if err != nil {
			return nil, fmt.Errorf("parsing time string: %w", err)
		}
		apiModels = append(apiModels, apiModel)
	}
	return apiModels, nil
}
