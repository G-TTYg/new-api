package openai

import (
	"encoding/json"
	"testing"

	"github.com/QuantumNous/new-api/dto"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConvertOpenAIResponsesRequestPreservesReasoningModeContextAndParsesMaxSuffix(t *testing.T) {
	t.Parallel()

	info := &relaycommon.RelayInfo{}
	got, err := (&Adaptor{}).ConvertOpenAIResponsesRequest(nil, info, dto.OpenAIResponsesRequest{
		Model: "gpt-5.6-sol-max",
		Reasoning: &dto.Reasoning{
			Mode:    json.RawMessage(`"pro"`),
			Context: json.RawMessage(`"all_turns"`),
		},
	})
	require.NoError(t, err)

	req, ok := got.(dto.OpenAIResponsesRequest)
	require.True(t, ok)
	assert.Equal(t, "gpt-5.6-sol", req.Model)
	require.NotNil(t, req.Reasoning)
	assert.Equal(t, "max", req.Reasoning.Effort)
	assert.Equal(t, "pro", req.Reasoning.Mode)
	assert.Equal(t, "all_turns", req.Reasoning.Context)
	assert.Equal(t, "max", info.ReasoningEffort)
}
