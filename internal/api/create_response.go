package api

import (
	"bytes"
	"cli-chat/internal/client"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func CreateResponse(client client.Client, model string, input string) (response, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/responses"

	payload := openAiPayload{
		Model:  model,
		Input:  input,
		Stream: false,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return response{}, fmt.Errorf("Error marshalling json: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return response{}, fmt.Errorf("Error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	res, err := client.HttpClient.Do(req)
	if err != nil {
		return response{}, fmt.Errorf("Error sending request: %w", err)
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return response{}, fmt.Errorf("Error reading request body: %w", err)
	}

	resp := response{}

	err = json.Unmarshal(data, &resp)
	if err != nil {
		return response{}, fmt.Errorf("Error unmarshalling json: %w", err)
	}

	return resp, nil
}
