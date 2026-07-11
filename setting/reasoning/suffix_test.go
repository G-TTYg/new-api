package reasoning

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseOpenAIReasoningEffortFromModelSuffixSupportsMax(t *testing.T) {
	effort, model := ParseOpenAIReasoningEffortFromModelSuffix("gpt-5.6-sol-max")

	require.Equal(t, "max", effort)
	require.Equal(t, "gpt-5.6-sol", model)
}
