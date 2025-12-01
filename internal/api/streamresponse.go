package api

import (
	"bufio"
	"bytes"
	"cli-chat/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Stream struct {
	Body   io.ReadCloser
	Cancel func()
}

type Delta struct {
	Type           string `json:"type"`
	SequenceNumber int    `json:"sequence_number"`
	ItemID         string `json:"item_id"`
	OutputIndex    int    `json:"output_index"`
	ContentIndex   int    `json:"content_index"`
	Delta          string `json:"delta"`
	Logprobs       []any  `json:"logprobs"`
	Obfuscation    string `json:"obfuscation"`
}

func CreateStreamResponse(ctx context.Context, cfg *config.Config, input string) (chan string, error) {
	payload := openAiPayload{
		Model:  cfg.AppSettings.Model,
		Input:  input,
		Stream: true,
	}
	stream, err := doStreamRequest(ctx, cfg, payload)
	if err != nil {
		return nil, fmt.Errorf("making stream request: %w", err)
	}

	outCh := make(chan string, 64)

	go func() {
		defer close(outCh)
		defer stream.Cancel()

		reader := bufio.NewReader(stream.Body)
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}

			line, err := reader.ReadBytes('\n')
			if err != nil {
			}
			if len(line) > 0 {
				if bytes.HasPrefix(line, []byte("data:")) {
					trimmed := bytes.TrimSpace(bytes.TrimPrefix(line, []byte("data:")))
					var delta Delta
					err = json.Unmarshal(trimmed, &delta)
					outCh <- delta.Delta
				}
			}
			if err != nil {
				if err != io.EOF {
					log.Printf("non-EOF error while reading from stream: %v", err)
					return
				}
				return
			}
		}
	}()
	return outCh, nil
}

func doStreamRequest(ctx context.Context, cfg *config.Config, payload openAiPayload) (*Stream, error) {
	apiKey := os.Getenv(cfg.AppSettings.Provider.Key)
	if apiKey == "" {
		return nil, fmt.Errorf("%s not set", cfg.AppSettings.Provider.Key)
	}
	url := cfg.AppSettings.Provider.BaseURL + "responses"
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshalling json: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(b))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")

	res, err := cfg.Client.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("doing request to %s: %w", url, err)
	}
	if res.StatusCode != http.StatusOK {
		defer res.Body.Close()
		msg, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("bad status %s: %s", res.Status, string(msg))
	}
	return &Stream{
		Body:   res.Body,
		Cancel: func() { _ = res.Body.Close() },
	}, nil
}
