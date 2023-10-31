package stage

import "github.com/upbound/upjet/pkg/config"

const shortGroup = "stage"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_stage_authenticator_duo", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AuthenticatorDuo"
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_sms", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AuthenticatorSMS"
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_static", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AuthenticatorStatic"
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_totp", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AuthenticatorTOTP"
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_validate", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AuthenticatorValidate"
	})
	p.AddResourceConfigurator("authentik_stage_authenticator_webauthn", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "AuthenticatorWebAuthn"
	})
	p.AddResourceConfigurator("authentik_stage_captcha", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_stage_consent", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_stage_deny", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_stage_dummy", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_stage_email", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_stage_identification", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_stage_invitation", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_stage_password", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_stage_prompt", func(r *config.Resource) {
		r.ShortGroup = shortGroup
	})
	p.AddResourceConfigurator("authentik_stage_prompt_field", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "PromptField"
	})
	p.AddResourceConfigurator("authentik_stage_user_delete", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UserDelete"
	})
	p.AddResourceConfigurator("authentik_stage_user_login", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UserLogin"
	})
	p.AddResourceConfigurator("authentik_stage_user_logout", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UserLogout"
	})
	p.AddResourceConfigurator("authentik_stage_user_write", func(r *config.Resource) {
		r.ShortGroup = shortGroup
		r.Kind = "UserWrite"
	})

}
