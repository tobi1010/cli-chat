package commands

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"terminal-chat/internal/api"
	"terminal-chat/internal/apitypes"
	"terminal-chat/session"
)

var (
	exclude = []string{
		"gpt-4-", "gpt-3", "dall", "image", "tts", "realtime", "audio", "turbo", "text",
		"embedding", "preview", "omni", "transcribe", "davinci", "babbage", "whisper", "search", "sora",
	}

	dateRegex = regexp.MustCompile(`-\d{4}-\d{2}-\d{2}$`)
)

func CommandListModels(s *session.Session, args []string) error {
	ctx := context.Background()

	models, err := api.GetModels(ctx, s.Client, s.ProviderName)
	if err != nil {
		return fmt.Errorf("requesting models: %w", err)
	}

	models = filterRelevantKeepLatest(models)
	sort.Slice(models, func(i, j int) bool {
		return models[i].Created > models[j].Created
	})

	for _, model := range models {
		if model.Name != "" {
			fmt.Println(model.Name)
		} else {
			fmt.Println(baseName(model.ID))
		}
	}

	return nil
}

func filterRelevantKeepLatest(models []apitypes.Model) []apitypes.Model {
	bestByBase := make(map[string]apitypes.Model, len(models))

	for _, m := range models {
		id := strings.ToLower(m.ID)
		if isExcluded(id) {
			continue
		}

		base := baseName(id)

		cur, ok := bestByBase[base]
		if !ok || m.Created > (cur.Created) {
			bestByBase[base] = m
			continue
		}

		// Optional tie-break: if timestamps equal, prefer lexicographically larger ID
		if m.Created == cur.Created && id > strings.ToLower(cur.ID) {
			bestByBase[base] = m
		}
	}

	out := make([]apitypes.Model, 0, len(bestByBase))
	for _, m := range bestByBase {
		out = append(out, m)
	}
	return out
}

func isExcluded(id string) bool {
	for _, ex := range exclude {
		if strings.Contains(id, ex) {
			return true
		}
	}
	return false
}

func baseName(name string) string {
	name = strings.ToLower(name)
	name = strings.TrimSuffix(name, "-chat-latest")
	name = strings.TrimSuffix(name, "-latest")
	return dateRegex.ReplaceAllString(name, "")
}
