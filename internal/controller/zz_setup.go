/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/MacroPower/provider-authentik/internal/controller/authentik/application"
	blueprint "github.com/MacroPower/provider-authentik/internal/controller/authentik/blueprint"
	certificatekeypair "github.com/MacroPower/provider-authentik/internal/controller/authentik/certificatekeypair"
	eventrule "github.com/MacroPower/provider-authentik/internal/controller/authentik/eventrule"
	eventtransport "github.com/MacroPower/provider-authentik/internal/controller/authentik/eventtransport"
	flow "github.com/MacroPower/provider-authentik/internal/controller/authentik/flow"
	flowstagebinding "github.com/MacroPower/provider-authentik/internal/controller/authentik/flowstagebinding"
	outpost "github.com/MacroPower/provider-authentik/internal/controller/authentik/outpost"
	scopemapping "github.com/MacroPower/provider-authentik/internal/controller/authentik/scopemapping"
	serviceconnectionkubernetes "github.com/MacroPower/provider-authentik/internal/controller/authentik/serviceconnectionkubernetes"
	tenant "github.com/MacroPower/provider-authentik/internal/controller/authentik/tenant"
	group "github.com/MacroPower/provider-authentik/internal/controller/directory/group"
	sourceldap "github.com/MacroPower/provider-authentik/internal/controller/directory/sourceldap"
	sourceoauth "github.com/MacroPower/provider-authentik/internal/controller/directory/sourceoauth"
	sourceplex "github.com/MacroPower/provider-authentik/internal/controller/directory/sourceplex"
	sourcesaml "github.com/MacroPower/provider-authentik/internal/controller/directory/sourcesaml"
	user "github.com/MacroPower/provider-authentik/internal/controller/directory/user"
	binding "github.com/MacroPower/provider-authentik/internal/controller/policy/binding"
	dummy "github.com/MacroPower/provider-authentik/internal/controller/policy/dummy"
	eventmatcher "github.com/MacroPower/provider-authentik/internal/controller/policy/eventmatcher"
	expiry "github.com/MacroPower/provider-authentik/internal/controller/policy/expiry"
	expression "github.com/MacroPower/provider-authentik/internal/controller/policy/expression"
	password "github.com/MacroPower/provider-authentik/internal/controller/policy/password"
	reputation "github.com/MacroPower/provider-authentik/internal/controller/policy/reputation"
	ldap "github.com/MacroPower/provider-authentik/internal/controller/propertymapping/ldap"
	notification "github.com/MacroPower/provider-authentik/internal/controller/propertymapping/notification"
	saml "github.com/MacroPower/provider-authentik/internal/controller/propertymapping/saml"
	scim "github.com/MacroPower/provider-authentik/internal/controller/propertymapping/scim"
	ldapprovider "github.com/MacroPower/provider-authentik/internal/controller/provider/ldap"
	oauth2 "github.com/MacroPower/provider-authentik/internal/controller/provider/oauth2"
	proxy "github.com/MacroPower/provider-authentik/internal/controller/provider/proxy"
	radius "github.com/MacroPower/provider-authentik/internal/controller/provider/radius"
	samlprovider "github.com/MacroPower/provider-authentik/internal/controller/provider/saml"
	scimprovider "github.com/MacroPower/provider-authentik/internal/controller/provider/scim"
	providerconfig "github.com/MacroPower/provider-authentik/internal/controller/providerconfig"
	authenticatorduo "github.com/MacroPower/provider-authentik/internal/controller/stage/authenticatorduo"
	authenticatorsms "github.com/MacroPower/provider-authentik/internal/controller/stage/authenticatorsms"
	authenticatorstatic "github.com/MacroPower/provider-authentik/internal/controller/stage/authenticatorstatic"
	authenticatortotp "github.com/MacroPower/provider-authentik/internal/controller/stage/authenticatortotp"
	authenticatorvalidate "github.com/MacroPower/provider-authentik/internal/controller/stage/authenticatorvalidate"
	authenticatorwebauthn "github.com/MacroPower/provider-authentik/internal/controller/stage/authenticatorwebauthn"
	captcha "github.com/MacroPower/provider-authentik/internal/controller/stage/captcha"
	consent "github.com/MacroPower/provider-authentik/internal/controller/stage/consent"
	deny "github.com/MacroPower/provider-authentik/internal/controller/stage/deny"
	dummystage "github.com/MacroPower/provider-authentik/internal/controller/stage/dummy"
	email "github.com/MacroPower/provider-authentik/internal/controller/stage/email"
	identification "github.com/MacroPower/provider-authentik/internal/controller/stage/identification"
	invitation "github.com/MacroPower/provider-authentik/internal/controller/stage/invitation"
	passwordstage "github.com/MacroPower/provider-authentik/internal/controller/stage/password"
	prompt "github.com/MacroPower/provider-authentik/internal/controller/stage/prompt"
	promptfield "github.com/MacroPower/provider-authentik/internal/controller/stage/promptfield"
	userdelete "github.com/MacroPower/provider-authentik/internal/controller/stage/userdelete"
	userlogin "github.com/MacroPower/provider-authentik/internal/controller/stage/userlogin"
	userlogout "github.com/MacroPower/provider-authentik/internal/controller/stage/userlogout"
	userwrite "github.com/MacroPower/provider-authentik/internal/controller/stage/userwrite"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		blueprint.Setup,
		certificatekeypair.Setup,
		eventrule.Setup,
		eventtransport.Setup,
		flow.Setup,
		flowstagebinding.Setup,
		outpost.Setup,
		scopemapping.Setup,
		serviceconnectionkubernetes.Setup,
		tenant.Setup,
		group.Setup,
		sourceldap.Setup,
		sourceoauth.Setup,
		sourceplex.Setup,
		sourcesaml.Setup,
		user.Setup,
		binding.Setup,
		dummy.Setup,
		eventmatcher.Setup,
		expiry.Setup,
		expression.Setup,
		password.Setup,
		reputation.Setup,
		ldap.Setup,
		notification.Setup,
		saml.Setup,
		scim.Setup,
		ldapprovider.Setup,
		oauth2.Setup,
		proxy.Setup,
		radius.Setup,
		samlprovider.Setup,
		scimprovider.Setup,
		providerconfig.Setup,
		authenticatorduo.Setup,
		authenticatorsms.Setup,
		authenticatorstatic.Setup,
		authenticatortotp.Setup,
		authenticatorvalidate.Setup,
		authenticatorwebauthn.Setup,
		captcha.Setup,
		consent.Setup,
		deny.Setup,
		dummystage.Setup,
		email.Setup,
		identification.Setup,
		invitation.Setup,
		passwordstage.Setup,
		prompt.Setup,
		promptfield.Setup,
		userdelete.Setup,
		userlogin.Setup,
		userlogout.Setup,
		userwrite.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
