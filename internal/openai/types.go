package openai

type openAiPayload struct {
	Model  string `json:"model"`
	Input  string `json:"input"`
	Stream bool   `json:"stream"`
}

type Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	OwnedBy string `json:"owned_by"`
}
