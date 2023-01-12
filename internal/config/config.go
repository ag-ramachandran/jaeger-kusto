package config

import (
	"errors"
	"os"
	"strings"

	"github.com/hashicorp/go-hclog"
	"github.com/spf13/viper"
)

// TODO : https://github.com/nicolastakashi/jaeger-redisearch/blob/main/internal/model/config.go

// KustoConfig contains AzureAD service principal and Kusto cluster configs
type KustoConfig struct {
	AuthType string `json:"auth_type"`      // Can be app or mi
	AppId    string `json:"application_id"` // Client Id for the AAD App
	/*If application auth is being used , then the following are relevant else they are not*/
	// Made changes to match these with configurations used in OpenTelemetry exporter
	AppKey          string `json:"application_key"`   // The client secret for the client
	TenantId        string `json:"tenant_id"`         // The tenant
	Endpoint        string `json:"cluster_uri"`       // Kusto cluster uri
	Database        string `json:"db_name"`           // database to query
	TracesTableName string `json:"traces_table_name"` // raw traces table
	SinkType        string `json:"sink_type"`         // can be OTEL, Jaeger (may be add additional backend in the future)

}

// ParseKustoConfig reads file at path and returns instance of KustoConfig or error
func ParseKustoConfig(logger hclog.Logger, configPath string) (*KustoConfig, error) {
	/*Initialize Viper to parse the config*/
	v := viper.New()
	v.AutomaticEnv()
	// can be passed in from env as well
	v.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	v.SetConfigType("json")
	/*Validate the file exists and can be read*/
	if configPath != "" {
		v.SetConfigFile(configPath)
		err := v.ReadInConfig()
		if err != nil {
			logger.Error("failed to parse kusto configuration file", "err", err)
			os.Exit(1)
		}
	}
	// Set any defaults we have. In this case the only default is the managed identity. The defaults are what is used in the OTEL Collector
	v.SetDefault("auth_type", "mi")
	v.SetDefault("traces_table_name", "OTELTraces")
	v.SetDefault("traces_table_name", "oteldb")
	v.SetDefault("sink_type", "otel")

	kc := &KustoConfig{}
	// mandatory to have the cluster url and database name
	kc.Endpoint = v.GetString("cluster_uri")
	// The common fields be in MI or App Auth
	kc.AuthType = v.GetString("auth_type")
	kc.Database = v.GetString("db_name")
	kc.TracesTableName = v.GetString("traces_table_name")
	kc.AppId = v.GetString("application_id")
	kc.Endpoint = v.GetString("cluster_uri")
	kc.SinkType = v.GetString("sink_type")

	// For app based auth only
	if kc.AuthType == "app" {
		// else parse the appid and app key
		kc.AppKey = v.GetString("application_key")
		kc.TenantId = v.GetString("tenant_id")
	}
	// Validate and return the config
	return Validate(kc)
}

// Validate returns error if any of required fields missing
func Validate(kc *KustoConfig) (*KustoConfig, error) {
	if kc.Database == "" {
		return nil, errors.New("missing database in kusto configuration")
	}
	if kc.Endpoint == "" {
		return nil, errors.New("missing cluster uri in kusto configuration")
	}
	if kc.AuthType == "auth" && (kc.AppId == "" || kc.AppKey == "" || kc.TenantId == "") {
		return nil, errors.New("missing client configuration (AppId, AppKey, TenantId) for kusto")
	}
	return kc, nil
}
