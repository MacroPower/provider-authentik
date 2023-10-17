package directory

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "directory"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_user", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_group", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		r.References["users"] = config.Reference{
			Type:      "User",
			Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("id",true)`,
		}
	})
}
