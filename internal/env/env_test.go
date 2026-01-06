package env

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestFindApiKeyCandidates(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "1234")
	t.Setenv("ANTHROPIC_API_KEY", "1234")
	keys := FindApiKeyCandidates()
	require.Contains(t, keys, "OPENAI_API_KEY")
	require.Contains(t, keys, "ANTHROPIC_API_KEY")
}
