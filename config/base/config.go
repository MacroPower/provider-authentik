package base

import "github.com/upbound/upjet/pkg/config"

const shortGroup = ""

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_application", func(r *config.Resource) {
		r.ShortGroup = shortGroup

		// // TODO: Support generic references
		// // https://github.com/crossplane/upjet/issues/95
		// r.References["protocol_provider"] = config.Reference{
		// 	Type:      "github.com/MacroPower/provider-authentik/apis/provider/v1alpha1.Proxy",
		// 	Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("id",true)`,
		// }
	})
	p.AddResourceConfigurator("authentik_outpost", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_service_connection_kubernetes", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ServiceConnectionKubernetes"
	})
	p.AddResourceConfigurator("authentik_blueprint", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_flow", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_flow_stage_binding", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "FlowStageBinding"
	})
	p.AddResourceConfigurator("authentik_scope_mapping", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "ScopeMapping"
	})
	p.AddResourceConfigurator("authentik_certificate_key_pair", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "CertificateKeyPair"
	})
	p.AddResourceConfigurator("authentik_tenant", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_event_rule", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EventRule"
	})
	p.AddResourceConfigurator("authentik_event_transport", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "EventTransport"
	})
}
