package openai

import (
	"cli-chat/internal/apitypes"
	"cli-chat/internal/client"
	"cli-chat/internal/env"
	"cli-chat/providers"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetModels(ctx context.Context, client *client.Client, provider providers.Provider) ([]apitypes.Model, error) {
	apiKey, err := env.ResolveAPIKey(provider.Key)
	if err != nil {
		return nil, fmt.Errorf("resloving API key for %s: %w", provider.Name, err)
	}
	url := provider.BaseURL + "models"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

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
	var modelsRes ModelsResponse
	if err = json.Unmarshal(data, &modelsRes); err != nil {
		return nil, fmt.Errorf("unmarshalling json: %w", err)
	}
	models := modelsRes.Data
	apiModels := []apitypes.Model{}
	for _, m := range models {
		apiModels = append(apiModels, toApiModel(m))
	}

	return apiModels, nil
}
