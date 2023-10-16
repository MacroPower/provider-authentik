package flow

import "github.com/upbound/upjet/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_flow", func(r *config.Resource) {
		r.ShortGroup = "flow"
	})
}
