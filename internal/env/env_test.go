package env

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindApiKeyCandidates(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "1234")
	keys := FindApiKeyCandidate()
	require.Equal(t, keys, []string{"OPENAI_API_KEY"})
}
