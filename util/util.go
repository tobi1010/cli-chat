package util

import (
	"cli-chat/internal/api"
	"strings"
)

func IsChatModel(model api.Model) bool {

	exclude := []string{
		"embedding", "moderation", "tts", "audio", "realtime",
		"image", "transcribe", "search", "sora", "diarize",
	}
	for _, ex := range exclude {
		if strings.Contains(model.ID, ex) {
			return false
		}
	}
	legacy := []string{
		"gpt-3.5", "davinci", "babbage",
		"gpt-4-0", "gpt-4-0613", "gpt-4-1106", "gpt-4-0125", "gpt-4-turbo-preview",
	}
	for _, leg := range legacy {
		if strings.Contains(model.ID, leg) {
			return false
		}
	}
	include := []string{
		// Families
		"gpt-5", "gpt-4o", "gpt-4.1", "o1", "o3",
		// Exact singleton IDs
		"gpt-5.1-chat-latest", "chatgpt-4o-latest",
	}
	for _, in := range include {
		if strings.HasPrefix(model.ID, in) || model.ID == in {
			return true
		}
	}
	return false
}
