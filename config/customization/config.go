package customization

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "customization"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_scope_mapping", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ScopeMapping"
	})
}
