package anthropic

import "cli-chat/internal/llm"

func ToApiModel(model Model) llm.Model {
	return llm.Model{ID: model.ID, Created: model.Created, OwnedBy: model.OwnedBy}
}
