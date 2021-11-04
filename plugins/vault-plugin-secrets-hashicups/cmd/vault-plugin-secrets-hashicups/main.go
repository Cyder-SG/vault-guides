package main

// Reference: https://learn.hashicorp.com/tutorials/vault/custom-secrets-engine-build?in=vault/custom-secrets-engine

import (
	"os"

	"github.com/hashicorp/go-hclog"
	hashicups "github.com/hashicorp/vault-guides/plugins/vault-plugin-secrets-hashicups"
	"github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/sdk/plugin"
)

func main() {
	apiClientMeta := &api.PluginAPIClientMeta{}
	flags := apiClientMeta.FlagSet()
	flags.Parse(os.Args[1:])

	tlsConfig := apiClientMeta.GetTLSConfig()
	tlsProviderFunc := api.VaultPluginTLSProvider(tlsConfig)

	err := plugin.Serve(&plugin.ServeOpts{
		BackendFactoryFunc: hashicups.CyderFactory,
		TLSProviderFunc:    tlsProviderFunc,
	})
	if err != nil {
		logger := hclog.New(&hclog.LoggerOptions{})

		logger.Error("plugin shutting down", "error", err)
		os.Exit(1)
	}
}
