package main

import (
	"cli-chat/config"
	"cli-chat/internal/api"
	"context"
	"fmt"
)

func commandListModels(cfg *config.Config, args []string) error {
	ctx := context.Background()
	models, err := api.GetModels(ctx, cfg)
	if err != nil {
		return fmt.Errorf("requesting models: %w", err)
	}
	for _, model := range models {
		if isChatModel(model) {
			fmt.Println(model.ID)
		}
	}
	return nil
}
