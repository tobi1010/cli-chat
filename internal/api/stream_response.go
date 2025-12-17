package api

import (
	"bufio"
	"bytes"
	"cli-chat/session"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Stream struct {
	Body   io.ReadCloser
	Cancel func()
}

type EventHeader struct {
	// for peeking at event header.type
	Type string `json:"type"`
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

type ResponseCompleted struct {
	Type           string         `json:"type"`
	SequenceNumber int            `json:"sequence_number"`
	Response       OpenAiResponse `json:"response"`
}

func CreateStreamResponse(ctx context.Context, s *session.Session, input string) (chan string, chan OpenAiResponse, chan error, error) {
	payload := openAiPayload{
		Model:  s.Provider.Model,
		Input:  input,
		Stream: true,
	}
	stream, err := doStreamRequest(ctx, s, payload)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("making stream request: %w", err)
	}

	outCh := make(chan string, 64)
	fullCh := make(chan OpenAiResponse)
	errCh := make(chan error)

	go func() {
		defer close(outCh)
		defer stream.Cancel()
		defer close(fullCh)
		defer close(errCh)

		reader := bufio.NewReader(stream.Body)
		for {
			select {
			case <-ctx.Done():
				errCh <- ctx.Err()
				return
			default:
			}

			line, err := reader.ReadBytes('\n')
			if err != nil {
			}
			if len(line) > 0 {
				if bytes.HasPrefix(line, []byte("data:")) {
					payload := bytes.TrimSpace(bytes.TrimPrefix(line, []byte("data:")))
					var eventHeader EventHeader
					if err = json.Unmarshal(payload, &eventHeader); err != nil {
						errCh <- fmt.Errorf("unmarshalling json: %w", err)
					}

					switch eventHeader.Type {
					case "response.output_text.delta":
						var delta Delta
						if err = json.Unmarshal(payload, &delta); err != nil {
							errCh <- fmt.Errorf("unmarshalling json: %w", err)
							return
						}
						outCh <- delta.Delta

					case "response.completed":
						var resCompletedEvent ResponseCompleted
						if err = json.Unmarshal(payload, &resCompletedEvent); err != nil {
							errCh <- fmt.Errorf("unmarshalling json: %w", err)
							return
						}
						fullCh <- resCompletedEvent.Response
						return
					default:
					}
				}
			}
			if err != nil {
				if err != io.EOF {
					errCh <- fmt.Errorf("non-EOF error while reading from stream: %v", err)
					return
				}
				return
			}
		}
	}()
	return outCh, fullCh, errCh, nil
}

func doStreamRequest(ctx context.Context, s *session.Session, payload openAiPayload) (*Stream, error) {
	apiKey := os.Getenv(s.Provider.Key)
	if apiKey == "" {
		return nil, fmt.Errorf("%s not set", s.Provider.Key)
	}
	url := s.Provider.BaseURL + "responses"
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

	res, err := s.Client.HttpClient.Do(req)
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
