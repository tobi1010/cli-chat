package openai

import (
	"cli-chat/internal/llm"
)

func toApiModel(model Model) llm.Model {
	return llm.Model{ID: model.ID, Created: model.Created, OwnedBy: model.OwnedBy}
}
func toApiResponse(response OpenAiResponse) llm.Response {
	return llm.Response{
		Model: response.Model,
	}
}
