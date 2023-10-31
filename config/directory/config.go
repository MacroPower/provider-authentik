package directory

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "directory"

var flowUUIDRef = config.Reference{
	Type:      "github.com/MacroPower/provider-authentik/apis/authentik/v1alpha1.Flow",
	Extractor: `github.com/upbound/upjet/pkg/resource.ExtractParamPath("uuid",true)`,
}

// Configure configures the directory provider.
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
	p.AddResourceConfigurator("authentik_source_ldap", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceLDAP"
	})
	p.AddResourceConfigurator("authentik_source_oauth", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceOAuth"

		r.References["authentication_flow"] = flowUUIDRef
		r.References["enrollment_flow"] = flowUUIDRef
	})
	p.AddResourceConfigurator("authentik_source_plex", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourcePlex"

		r.References["authentication_flow"] = flowUUIDRef
		r.References["enrollment_flow"] = flowUUIDRef
	})
	p.AddResourceConfigurator("authentik_source_saml", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "SourceSAML"

		r.References["authentication_flow"] = flowUUIDRef
		r.References["enrollment_flow"] = flowUUIDRef
		r.References["pre_authentication_flow"] = flowUUIDRef
	})
}
