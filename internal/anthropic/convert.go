package anthropic

import (
	"cli-chat/internal/apitypes"
	"time"
)

func toApiModel(model Model) (apitypes.Model, error) {
	t, err := time.Parse(time.RFC3339, model.CreatedAt)
	if err != nil {
		return apitypes.Model{}, err
	}
	secs := t.Unix()
	return apitypes.Model{
		ID:      model.ID,
		Created: secs,
		Name:    model.DisplayName,
	}, nil

}

func toApiResponse(res AnthropicResponse) apitypes.Response {
	return apitypes.Response{}
}
