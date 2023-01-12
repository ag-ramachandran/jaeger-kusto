package main

import (
	"flag"
	"kusto-jaeger-plugin/internal/config"
	"os"

	"github.com/Azure/azure-kusto-go/kusto"

	"github.com/hashicorp/go-hclog"
)

const (
	loggerName = "kusto-jaeger-plugin"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "config", "", "A path to the plugin's configuration file")
	flag.Parse()
	/*
		Get the logger for the application
	*/
	logger := hclog.New(&hclog.LoggerOptions{
		Name:       loggerName,
		Level:      hclog.Warn, // Warn and up only
		JSONFormat: true,
	})

	/*
		Parse the configs passed in
	*/
	kc, err := config.ParseKustoConfig(logger, configPath)

	if err != nil {
		logger.Error("error parsing kusto connection config", "error", err)
		os.Exit(1)
	}

	var kcsb *kusto.ConnectionStringBuilder

	if kc.AuthType == "mi" {
		kcsb = kusto.NewConnectionStringBuilder(kc.Endpoint).WithSystemManagedIdentity()
	} else {
		logger.Info("Connecting to Kusto cluster using application id credentials")
		kcsb = kusto.NewConnectionStringBuilder(kc.Endpoint).WithAadAppKey(kc.AppId, kc.AppKey, kc.TenantId)
	}

	c, cerr := kusto.New(kcsb)

	if cerr != nil {
		logger.Error("error connecting to kusto", "error", cerr)
		os.Exit(2)
	}
	defer c.Close()
}
