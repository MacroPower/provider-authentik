package application

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_application", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "application"
	})
	p.AddResourceConfigurator("authentik_provider_oauth2", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource, which would be "github"
		r.ShortGroup = "application"
		r.Kind = "ProviderOAuth2"

		r.References["authorization_flow"] = config.Reference{
			Type:      "github.com/MacroPower/provider-authentik/apis/flow/v1alpha1.Flow",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("uuid",true)`,
		}
	})
}
