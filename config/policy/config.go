package policy

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "policy"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_policy_binding", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Binding"
	})
	p.AddResourceConfigurator("authentik_policy_dummy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Dummy"
	})
	p.AddResourceConfigurator("authentik_policy_event_matcher", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EventMatcher"
	})
	p.AddResourceConfigurator("authentik_policy_expiry", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Expiry"
	})
	p.AddResourceConfigurator("authentik_policy_expression", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Expression"
	})
	p.AddResourceConfigurator("authentik_policy_password", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Password"
	})
	p.AddResourceConfigurator("authentik_policy_reputation", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "Reputation"
	})
}
