package commands

import (
	"cli-chat/internal/api"
	"cli-chat/session"
	"cli-chat/util"
	"context"
	"fmt"
)

func CommandListModels(s *session.Session, args []string) error {
	ctx := context.Background()
	models, err := api.GetModels(ctx, s)
	if err != nil {
		return fmt.Errorf("requesting models: %w", err)
	}
	for _, model := range models {
		if util.IsChatModel(model) {
			fmt.Println(model.ID)
		}
	}
	return nil
}
