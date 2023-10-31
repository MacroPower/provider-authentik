/*
Copyright 2022 Upbound Inc.
*/

package config

import "github.com/upbound/upjet/pkg/config"

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"authentik_application":                   config.ParameterAsIdentifier("slug"),
	"authentik_outpost":                       config.IdentifierFromProvider,
	"authentik_service_connection_kubernetes": config.IdentifierFromProvider,
	"authentik_blueprint":                     config.IdentifierFromProvider,
	"authentik_flow":                          config.ParameterAsIdentifier("slug"),
	"authentik_flow_stage_binding":            config.IdentifierFromProvider,
	"authentik_scope_mapping":                 config.IdentifierFromProvider,
	"authentik_certificate_key_pair":          config.IdentifierFromProvider,
	"authentik_tenant":                        config.IdentifierFromProvider,
	"authentik_event_rule":                    config.IdentifierFromProvider,
	"authentik_event_transport":               config.IdentifierFromProvider,

	"authentik_user":         config.IdentifierFromProvider,
	"authentik_group":        config.IdentifierFromProvider,
	"authentik_source_ldap":  config.IdentifierFromProvider,
	"authentik_source_oauth": config.IdentifierFromProvider,
	"authentik_source_plex":  config.IdentifierFromProvider,
	"authentik_source_saml":  config.IdentifierFromProvider,

	"authentik_policy_binding":       config.IdentifierFromProvider,
	"authentik_policy_dummy":         config.IdentifierFromProvider,
	"authentik_policy_event_matcher": config.IdentifierFromProvider,
	"authentik_policy_expiry":        config.IdentifierFromProvider,
	"authentik_policy_expression":    config.IdentifierFromProvider,
	"authentik_policy_password":      config.IdentifierFromProvider,
	"authentik_policy_reputation":    config.IdentifierFromProvider,

	"authentik_property_mapping_ldap":         config.IdentifierFromProvider,
	"authentik_property_mapping_notification": config.IdentifierFromProvider,
	"authentik_property_mapping_saml":         config.IdentifierFromProvider,
	"authentik_property_mapping_scim":         config.IdentifierFromProvider,

	"authentik_provider_ldap":   config.IdentifierFromProvider,
	"authentik_provider_oauth2": config.IdentifierFromProvider,
	"authentik_provider_proxy":  config.IdentifierFromProvider,
	"authentik_provider_radius": config.IdentifierFromProvider,
	"authentik_provider_saml":   config.IdentifierFromProvider,
	"authentik_provider_scim":   config.IdentifierFromProvider,

	"authentik_stage_authenticator_duo":      config.IdentifierFromProvider,
	"authentik_stage_authenticator_sms":      config.IdentifierFromProvider,
	"authentik_stage_authenticator_static":   config.IdentifierFromProvider,
	"authentik_stage_authenticator_totp":     config.IdentifierFromProvider,
	"authentik_stage_authenticator_validate": config.IdentifierFromProvider,
	"authentik_stage_authenticator_webauthn": config.IdentifierFromProvider,
	"authentik_stage_captcha":                config.IdentifierFromProvider,
	"authentik_stage_consent":                config.IdentifierFromProvider,
	"authentik_stage_deny":                   config.IdentifierFromProvider,
	"authentik_stage_dummy":                  config.IdentifierFromProvider,
	"authentik_stage_email":                  config.IdentifierFromProvider,
	"authentik_stage_identification":         config.IdentifierFromProvider,
	"authentik_stage_invitation":             config.IdentifierFromProvider,
	"authentik_stage_password":               config.IdentifierFromProvider,
	"authentik_stage_prompt":                 config.IdentifierFromProvider,
	"authentik_stage_prompt_field":           config.IdentifierFromProvider,
	"authentik_stage_user_delete":            config.IdentifierFromProvider,
	"authentik_stage_user_login":             config.IdentifierFromProvider,
	"authentik_stage_user_logout":            config.IdentifierFromProvider,
	"authentik_stage_user_write":             config.IdentifierFromProvider,
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
