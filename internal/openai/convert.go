package openai

import "terminal-chat/internal/apitypes"

func toApiModel(model Model) apitypes.Model {
	return apitypes.Model{
		ID:      model.ID,
		Created: model.Created,
		Name:    model.ID,
	}

}
func toApiResponse(res OpenAiResponse) apitypes.Response {
	return apitypes.Response{
		ID:    res.ID,
		Model: res.Model,
	}
}
