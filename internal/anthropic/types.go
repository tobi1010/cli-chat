package anthropic

type anthropicPayload struct {
	Model  string `json:"model"`
	Input  string `json:"input"`
	Stream bool   `json:"stream"`
}

type deltaEvent struct {
	Type           string   `json:"type"`
	SequenceNumber int      `json:"sequence_number"`
	Response       struct{} `json:"response"`
	Delta          string   `json:"delta"`
}

type anthropicResponse struct {
	ID                string `json:"id"`
	Object            string `json:"object"`
	CreatedAt         int    `json:"created_at"`
	Status            string `json:"status"`
	Error             any    `json:"error"`
	IncompleteDetails any    `json:"incomplete_details"`
	Instructions      any    `json:"instructions"`
	MaxOutputTokens   any    `json:"max_output_tokens"`
	Model             string `json:"model"`
	Output            []struct {
		Type    string `json:"type"`
		ID      string `json:"id"`
		Status  string `json:"status"`
		Role    string `json:"role"`
		Content []struct {
			Type        string `json:"type"`
			Text        string `json:"text"`
			Annotations []any  `json:"annotations"`
		} `json:"content"`
	} `json:"output"`
	ParallelToolCalls  bool `json:"parallel_tool_calls"`
	PreviousResponseID any  `json:"previous_response_id"`
	Reasoning          struct {
		Effort  any `json:"effort"`
		Summary any `json:"summary"`
	} `json:"reasoning"`
	Store       bool    `json:"store"`
	Temperature float64 `json:"temperature"`
	Text        struct {
		Format struct {
			Type string `json:"type"`
		} `json:"format"`
	} `json:"text"`
	ToolChoice string  `json:"tool_choice"`
	Tools      []any   `json:"tools"`
	TopP       float64 `json:"top_p"`
	Truncation string  `json:"truncation"`
	Usage      struct {
		InputTokens        int `json:"input_tokens"`
		InputTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
		} `json:"input_tokens_details"`
		OutputTokens        int `json:"output_tokens"`
		OutputTokensDetails struct {
			ReasoningTokens int `json:"reasoning_tokens"`
		} `json:"output_tokens_details"`
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
	User     any `json:"user"`
	Metadata struct {
	} `json:"metadata"`
}

type anthropicStreamResponse struct {
	Type           string `json:"type"`
	SequenceNumber int    `json:"sequence_number"`
	Response       struct {
		ID                string `json:"id"`
		Object            string `json:"object"`
		CreatedAt         int    `json:"created_at"`
		Status            string `json:"status"`
		Background        bool   `json:"background"`
		Error             any    `json:"error"`
		IncompleteDetails any    `json:"incomplete_details"`
		Instructions      any    `json:"instructions"`
		MaxOutputTokens   any    `json:"max_output_tokens"`
		MaxToolCalls      any    `json:"max_tool_calls"`
		Model             string `json:"model"`
		Output            []struct {
			ID      string `json:"id"`
			Type    string `json:"type"`
			Role    string `json:"role"`
			Status  string `json:"status"`
			Content []struct {
				Type        string `json:"type"`
				Text        string `json:"text"`
				Annotations any    `json:"annotations"`
				Logprobs    any    `json:"logprobs"`
			} `json:"content"`
		} `json:"output"`
		ParallelToolCalls    bool `json:"parallel_tool_calls"`
		PreviousResponseID   any  `json:"previous_response_id"`
		PromptCacheKey       any  `json:"prompt_cache_key"`
		PromptCacheRetention any  `json:"prompt_cache_retention"`
		Reasoning            struct {
			Effort  string `json:"effort"`
			Summary any    `json:"summary"`
		} `json:"reasoning"`
		SafetyIdentifier any     `json:"safety_identifier"`
		ServiceTier      string  `json:"service_tier"`
		Store            bool    `json:"store"`
		Temperature      float64 `json:"temperature"`
		Text             struct {
			Format struct {
				Type string `json:"type"`
			} `json:"format"`
			Verbosity string `json:"verbosity"`
		} `json:"text"`
		ToolChoice  string  `json:"tool_choice"`
		Tools       []any   `json:"tools"`
		TopLogprobs int     `json:"top_logprobs"`
		TopP        float64 `json:"top_p"`
		Truncation  string  `json:"truncation"`
		Usage       any     `json:"usage"`
		User        any     `json:"user"`
		Metadata    struct {
		} `json:"metadata"`
	} `json:"response"`
}
