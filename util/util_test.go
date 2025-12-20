package util

import (
	"cli-chat/internal/api"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsChatModel(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want bool
	}{
		// include: prefixes
		{"gpt-5 prefix", "gpt-5", true},
		{"gpt-5 variant", "gpt-5-mini", true},
		{"gpt-4o prefix", "gpt-4o-2024-08-06", true},
		{"gpt-4.1 prefix", "gpt-4.1-mini", true},
		{"o1 prefix", "o1-mini", true},
		{"o3 prefix", "o3-mini", true},

		// include: exact singletons
		{"gpt-5.1-chat-latest exact", "gpt-5.1-chat-latest", true},
		{"chatgpt-4o-latest exact", "chatgpt-4o-latest", true},

		// exclude substrings
		{"embedding excluded", "text-embedding-3-large", false},
		{"moderation excluded", "omni-moderation-latest", false},
		{"tts excluded", "gpt-4o-mini-tts", false},
		{"audio excluded", "gpt-4o-audio-preview", false},
		{"realtime excluded", "gpt-4o-realtime-preview", false},
		{"image excluded", "gpt-image-1", false},
		{"transcribe excluded", "gpt-4o-transcribe", false},
		{"search excluded", "gpt-4o-search-preview", false},
		{"sora excluded", "sora-1", false},
		{"diarize excluded", "gpt-4o-diarize", false},

		// legacy excluded
		{"gpt-3.5 legacy", "gpt-3.5-turbo", false},
		{"davinci legacy", "text-davinci-003", false},
		{"babbage legacy", "babbage-002", false},
		{"gpt-4-0613 legacy", "gpt-4-0613", false},
		{"gpt-4-turbo-preview legacy", "gpt-4-turbo-preview", false},

		// unknown defaults
		{"unknown model", "some-new-model", false},

		// precedence: exclude wins even if included prefix matches
		{"exclude wins over include", "gpt-5-image", false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsChatModel(api.Model{ID: tc.id})
			require.Equal(t, tc.want, got)
		})
	}
}
