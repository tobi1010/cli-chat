package anthropic

import (
	"cli-chat/internal/apitypes"
)

func toApiModel(model Model) apitypes.Model {
	return apitypes.Model{
		ID:      model.ID,
		Created: model.Created,
		OwnedBy: model.OwnedBy,
	}

}

func toApiResponse(res anthropicResponse) apitypes.Response {
	return apitypes.Response{}
}
