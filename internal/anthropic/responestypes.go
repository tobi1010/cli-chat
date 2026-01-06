package anthropic

type AnthropicModelsResponse struct {
	Data    []Model `json:"data"`
	FirstId string  `json:"first_id"`
	HasMore bool    `json:"has_more"`
	LastId  string  `json:"last_id"`
}

type deltaEvent struct {
	Type           string   `json:"type"`
	SequenceNumber int      `json:"sequence_number"`
	Response       struct{} `json:"response"`
	Delta          string   `json:"delta"`
}

type AnthropicResponse struct {
	ID      string `json:"id"`
	Content []struct {
		Citations []struct {
			CitedText      string `json:"cited_text"`
			DocumentIndex  int    `json:"document_index"`
			DocumentTitle  string `json:"document_title"`
			EndCharIndex   int    `json:"end_char_index"`
			FileID         string `json:"file_id"`
			StartCharIndex int    `json:"start_char_index"`
			Type           string `json:"type"`
		} `json:"citations"`
		Text string `json:"text"`
		Type string `json:"type"`
	} `json:"content"`
	Model        string `json:"model"`
	Role         string `json:"role"`
	StopReason   string `json:"stop_reason"`
	StopSequence any    `json:"stop_sequence"`
	Type         string `json:"type"`
	Usage        struct {
		CacheCreation struct {
			Ephemeral1HInputTokens int `json:"ephemeral_1h_input_tokens"`
			Ephemeral5MInputTokens int `json:"ephemeral_5m_input_tokens"`
		} `json:"cache_creation"`
		CacheCreationInputTokens int `json:"cache_creation_input_tokens"`
		CacheReadInputTokens     int `json:"cache_read_input_tokens"`
		InputTokens              int `json:"input_tokens"`
		OutputTokens             int `json:"output_tokens"`
		ServerToolUse            struct {
			WebSearchRequests int `json:"web_search_requests"`
		} `json:"server_tool_use"`
		ServiceTier string `json:"service_tier"`
	} `json:"usage"`
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
