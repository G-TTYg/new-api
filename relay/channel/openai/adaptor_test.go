package openai

import (
	"testing"

	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/dto"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	relayconstant "github.com/QuantumNous/new-api/relay/constant"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAzureResponsesRequestURL(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name                  string
		baseURL               string
		relayMode             int
		apiVersion            string
		azureResponsesVersion string
		expected              string
	}{
		{
			name:       "default openai azure endpoint uses v1 without api-version",
			baseURL:    "https://example.openai.azure.com",
			relayMode:  relayconstant.RelayModeResponses,
			apiVersion: "2025-04-01-preview",
			expected:   "https://example.openai.azure.com/openai/v1/responses",
		},
		{
			name:                  "preview setting maps to v1 without api-version",
			baseURL:               "https://example.openai.azure.com",
			relayMode:             relayconstant.RelayModeResponses,
			azureResponsesVersion: "preview",
			expected:              "https://example.openai.azure.com/openai/v1/responses",
		},
		{
			name:                  "compact uses v1 compact path",
			baseURL:               "https://example.services.ai.azure.com",
			relayMode:             relayconstant.RelayModeResponsesCompact,
			azureResponsesVersion: "v1",
			expected:              "https://example.services.ai.azure.com/openai/v1/responses/compact",
		},
		{
			name:                  "base url with openai v1 is not duplicated",
			baseURL:               "https://example.openai.azure.com/openai/v1",
			relayMode:             relayconstant.RelayModeResponses,
			azureResponsesVersion: "v1",
			expected:              "https://example.openai.azure.com/openai/v1/responses",
		},
		{
			name:       "cognitive services default keeps dated api-version",
			baseURL:    "https://example.cognitiveservices.azure.com",
			relayMode:  relayconstant.RelayModeResponses,
			apiVersion: "2025-04-01-preview",
			expected:   "https://example.cognitiveservices.azure.com/openai/responses?api-version=2025-04-01-preview",
		},
		{
			name:                  "explicit dated version keeps dated api-version path",
			baseURL:               "https://example.openai.azure.com",
			relayMode:             relayconstant.RelayModeResponses,
			azureResponsesVersion: "2025-04-01-preview",
			expected:              "https://example.openai.azure.com/openai/responses?api-version=2025-04-01-preview",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			info := azureResponsesTestInfo(tt.baseURL, tt.relayMode)
			info.ApiVersion = tt.apiVersion
			info.ChannelOtherSettings.AzureResponsesVersion = tt.azureResponsesVersion

			got, err := (&Adaptor{}).GetRequestURL(info)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func azureResponsesTestInfo(baseURL string, relayMode int) *relaycommon.RelayInfo {
	return &relaycommon.RelayInfo{
		RelayMode:      relayMode,
		RequestURLPath: "/v1/responses",
		ChannelMeta: &relaycommon.ChannelMeta{
			ChannelType:       constant.ChannelTypeAzure,
			ChannelBaseUrl:    baseURL,
			UpstreamModelName: "gpt-5.6-sol",
			ChannelOtherSettings: dto.ChannelOtherSettings{
				AzureResponsesVersion: "",
			},
		},
	}
}
