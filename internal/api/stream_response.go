package api

import (
	"bufio"
	"bytes"
	"cli-chat/internal/client"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type deltaEvent struct {
	Type           string   `json:"type"`
	SequenceNumber int      `json:"sequence_number"`
	Response       struct{} `json:"response"`
	Delta          string   `json:"delta"`
}

func StreamResponse(client *client.Client, model, input string) error {
	apiKey := os.Getenv("OPENAI_API_KEY")
	url := "https://api.openai.com/v1/responses"

	payload := openAiPayload{
		Model:  model,
		Input:  input,
		Stream: true,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling json: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")

	res, err := client.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", res.Status)
	}

	reader := bufio.NewReader(res.Body)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		if len(line) < 6 || !bytes.HasPrefix(line, []byte("data: ")) {
			continue
		}
		data := bytes.TrimPrefix(line, []byte("data: "))
		trimmedData := bytes.TrimSpace(data)
		if bytes.Equal([]byte("[DONE]"), trimmedData) {
			fmt.Println("done")
			break
		}
		var streamDelta deltaEvent
		err = json.Unmarshal(trimmedData, &streamDelta)
		if err != nil {
			return fmt.Errorf("error unmashalling json: %w", err)
		}
		if streamDelta.Type == "response.output_text.delta" {
			fmt.Print(streamDelta.Delta)
			continue
		}

		// var fullResponse streamRespponse
		// err = json.Unmarshal(data, &fullResponse)
		// if err != nil {
		// 	return fmt.Errorf("error unmashalling json: %w", err)
		// }
		// for _, output := range fullResponse.Response.Output {
		// 	for _, content := range output.Content {
		// 		fmt.Println(content.Text)
		// 	}
		// }

	}
	fmt.Print("\n")

	return nil
}
