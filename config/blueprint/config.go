package blueprint

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "blueprint"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_blueprint", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
}
