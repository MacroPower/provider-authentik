/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/upbound/upjet/pkg/config"

	"github.com/MacroPower/provider-authentik/config/base"
	"github.com/MacroPower/provider-authentik/config/directory"
	"github.com/MacroPower/provider-authentik/config/policy"
	"github.com/MacroPower/provider-authentik/config/propertymapping"
	"github.com/MacroPower/provider-authentik/config/provider"
	"github.com/MacroPower/provider-authentik/config/stage"
)

const (
	resourcePrefix = "authentik"
	modulePath     = "github.com/MacroPower/provider-authentik"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("authentik.crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		base.Configure,
		directory.Configure,
		policy.Configure,
		propertymapping.Configure,
		provider.Configure,
		stage.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
