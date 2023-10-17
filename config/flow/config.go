package flow

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "flow"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_flow", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_flow_stage_binding", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "StageBinding"
	})
}
