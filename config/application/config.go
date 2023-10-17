package application

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_application", func(r *config.Resource) {
		r.ShortGroup = "application"

		r.References["protocol_provider"] = config.Reference{
			Type:      "github.com/MacroPower/provider-authentik/apis/provider/v1alpha1.Proxy",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
}
