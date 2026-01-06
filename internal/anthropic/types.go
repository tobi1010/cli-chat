package anthropic

type anthropicPayload struct {
	Model  string `json:"model"`
	Input  string `json:"input"`
	Stream bool   `json:"stream"`
}

type Model struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
}
