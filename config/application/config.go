package application

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "application"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_application", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		// TODO: Support generic references
		// https://github.com/crossplane/upjet/issues/95
		r.References["protocol_provider"] = config.Reference{
			Type:      "github.com/MacroPower/provider-authentik/apis/provider/v1alpha1.Proxy",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
	p.AddResourceConfigurator("authentik_outpost", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		r.References["protocol_provider"] = config.Reference{
			Type:      "github.com/MacroPower/provider-authentik/apis/provider/v1alpha1.Proxy",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
	p.AddResourceConfigurator("authentik_service_connection_kubernetes", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ServiceConnectionKubernetes"
	})
}
