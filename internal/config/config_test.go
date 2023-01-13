package config

import (
	"path/filepath"
	"testing"

	"github.com/hashicorp/go-hclog"
	"github.com/stretchr/testify/require"
	"github.com/tj/assert"
)

// ParseKustoConfig reads file at path and returns instance of KustoConfig or error
func TestParseKustoConfig(t *testing.T) {
	t.Parallel()
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       "jaeger-kusto-test",
		Level:      hclog.Warn, // Warn and up only
		JSONFormat: true,
	})

	tests := []struct {
		id             string
		configFilePath string
		expected       *KustoConfig
		errorMessage   string
	}{
		{
			id:             "app-auth-success",
			configFilePath: filepath.Join("testdata", "conf-test-app-success.json"),
			expected: &KustoConfig{
				AuthType:        "app",
				AppId:           "app-id",
				AppKey:          "app-key",
				TenantId:        "tenant-id",
				Endpoint:        "https://test-otel.dev.kusto.windows.net",
				Database:        "otel-db",
				TracesTableName: "otel-traces",
				SinkType:        "jaeger",
			},
			errorMessage: "",
		},
		{
			id:             "app-mi-success",
			configFilePath: filepath.Join("testdata", "conf-test-mi-success.json"),
			expected: &KustoConfig{
				AuthType:        "mi",
				AppId:           "app-id",
				Endpoint:        "https://test-otel.dev.kusto.windows.net",
				Database:        "otel-db",
				TracesTableName: "otel-traces",
				SinkType:        "otel",
			},
			errorMessage: "",
		},
	}
	// Run these tests
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
			got, err := ParseKustoConfig(logger, tt.configFilePath)
			require.NoError(t, err)
			require.NotNil(t, got)
			assert.Equal(t, tt.expected, got)
		})
	}
}
