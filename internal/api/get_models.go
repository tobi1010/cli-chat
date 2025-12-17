package api

import (
	"cli-chat/internal/auth"
	"cli-chat/session"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	OwnedBy string `json:"owned_by"`
}

type ModelsResponse struct {
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

func GetModels(ctx context.Context, s *session.Session) ([]Model, error) {
	apiKey, err := auth.ResolveAPIKey(s.Provider.Key)
	if err != nil {
		return nil, fmt.Errorf("resloving API key for %s: %w", s.Provider, err)
	}
	url := s.Provider.BaseURL + "models"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)

	res, err := s.Client.HttpClient.Do(req)
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

	return models, nil
}
