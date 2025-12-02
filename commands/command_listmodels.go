package commands

import (
	"cli-chat/config"
	"cli-chat/internal/api"
	"cli-chat/util"
	"context"
	"fmt"
)

func CommandListModels(cfg *config.Config, args []string) error {
	ctx := context.Background()
	models, err := api.GetModels(ctx, cfg)
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
