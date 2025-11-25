package api

import (
	"bufio"
	"bytes"
	"cli-chat/internal/client"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const key = "OPENAI_API_KEY"
const urlString = "https://api.openai.com/v1/responses"

func StreamDeltas(
	ctx context.Context,
	client *client.Client,
	model string,
	input string,
) (
	deltas <-chan string,
	done <-chan StreamResponse,
	errc <-chan error,
	cancel func(),
	err error,
) {
	fmt.Println("[dbg] streaming deltas")
	apiKey := os.Getenv(key)
	if apiKey == "" {
		return nil, nil, nil, nil, fmt.Errorf("%s not set", key)
	}
	fmt.Println("here1")

	payload := openAiPayload{
		Model:  model,
		Input:  input,
		Stream: true,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("marshalling JSON: %w", err)
	}
	url := urlString
	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("creating request %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")

	res, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("doing request: %w", err)
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		b, _ := io.ReadAll(res.Body)
		return nil, nil, nil, nil, fmt.Errorf("bad status: %s; body: %s", res.Status, string(b))
	}
	outCh := make(chan string, 64)
	fullCh := make(chan StreamResponse, 1)
	errCh := make(chan error, 1)

	cancel = func() {
		_ = res.Body.Close()
	}

	fmt.Println("[dbg] starting goroutine")
	go func() {
		fmt.Println("[dbg] running channels")
		defer close(outCh)
		defer close(fullCh)
		defer close(errCh)
		defer res.Body.Close()

		reader := bufio.NewReader(res.Body)
		for {
			line, err := reader.ReadBytes('\n')
			if err != nil {
				if err != io.EOF {
					errCh <- fmt.Errorf("reading stream: %w", err)
				}
				return
			}
			fmt.Println("[dbg] line:", string(line))
			if !bytes.HasPrefix(line, []byte("data: ")) {
				continue
			}
			data := bytes.TrimSpace(bytes.TrimPrefix(line, []byte("data: ")))
			if bytes.Equal(data, []byte("[DONE]")) {
				return
			}

			var deltaEv deltaEvent
			err = json.Unmarshal(data, &deltaEv)
			if err != nil {
				errCh <- fmt.Errorf("unmarshalling JSON: %w", err)
				return
			}
			if deltaEv.Type == "response.output_text.delta" {
				if deltaEv.Delta != "" {
					outCh <- deltaEv.Delta
				}
				continue
			}
			var fullResponse StreamResponse
			err = json.Unmarshal(data, &fullResponse)
			if err != nil {
				errCh <- fmt.Errorf("unmarshalling JSON: %w", err)
				return
			}
			if fullResponse.Response.ID != "" {
				fullCh <- fullResponse
				return
			}
		}
	}()
	return outCh, fullCh, errCh, cancel, nil
}

// func StreamResponse(client *client.Client, model, input string) error {
// 	apiKey := os.Getenv(key)
// 	url := urlString
//
// 	payload := openAiPayload{
// 		Model:  model,
// 		Input:  input,
// 		Stream: true,
// 	}
//
// 	body, err := json.Marshal(payload)
// 	if err != nil {
// 		return fmt.Errorf("error marshalling json: %w", err)
// 	}
//
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return fmt.Errorf("error creating request: %w", err)
// 	}
//
// 	req.Header.Set("Authorization", "Bearer "+apiKey)
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Accept", "text/event-stream")
//
// 	res, err := client.HttpClient.Do(req)
// 	if err != nil {
// 		return fmt.Errorf("error making request: %w", err)
// 	}
// 	defer res.Body.Close()
//
// 	if res.StatusCode != http.StatusOK {
// 		return fmt.Errorf("bad status: %s", res.Status)
// 	}
//
// 	reader := bufio.NewReader(res.Body)
// 	for {
// 		line, err := reader.ReadBytes('\n')
// 		if err != nil {
// 			break
// 		}
// 		if len(line) < 6 || !bytes.HasPrefix(line, []byte("data: ")) {
// 			continue
// 		}
// 		data := bytes.TrimPrefix(line, []byte("data: "))
// 		trimmedData := bytes.TrimSpace(data)
// 		if bytes.Equal([]byte("[DONE]"), trimmedData) {
// 			fmt.Println("done")
// 			break
// 		}
// 		var streamDelta deltaEvent
// 		err = json.Unmarshal(trimmedData, &streamDelta)
// 		if err != nil {
// 			return fmt.Errorf("error unmashalling json: %w", err)
// 		}
// 		if streamDelta.Type == "response.output_text.delta" {
// 			fmt.Print(streamDelta.Delta)
// 			continue
// 		}
//
// 		// var fullResponse streamRespponse
// 		// err = json.Unmarshal(data, &fullResponse)
// 		// if err != nil {
// 		// 	return fmt.Errorf("error unmashalling json: %w", err)
// 		// }
// 		// for _, output := range fullResponse.Response.Output {
// 		// 	for _, content := range output.Content {
// 		// 		fmt.Println(content.Text)
// 		// 	}
// 		// }
//
// 	}
// 	fmt.Print("\n")
//
// 	return nil
// }
