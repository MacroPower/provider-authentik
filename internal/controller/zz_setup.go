/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/MacroPower/provider-authentik/internal/controller/application/application"
	outpost "github.com/MacroPower/provider-authentik/internal/controller/application/outpost"
	serviceconnectionkubernetes "github.com/MacroPower/provider-authentik/internal/controller/application/serviceconnectionkubernetes"
	blueprint "github.com/MacroPower/provider-authentik/internal/controller/blueprint/blueprint"
	scopemapping "github.com/MacroPower/provider-authentik/internal/controller/customization/scopemapping"
	group "github.com/MacroPower/provider-authentik/internal/controller/directory/group"
	user "github.com/MacroPower/provider-authentik/internal/controller/directory/user"
	flow "github.com/MacroPower/provider-authentik/internal/controller/flow/flow"
	stagebinding "github.com/MacroPower/provider-authentik/internal/controller/flow/stagebinding"
	ldap "github.com/MacroPower/provider-authentik/internal/controller/provider/ldap"
	oauth2 "github.com/MacroPower/provider-authentik/internal/controller/provider/oauth2"
	proxy "github.com/MacroPower/provider-authentik/internal/controller/provider/proxy"
	radius "github.com/MacroPower/provider-authentik/internal/controller/provider/radius"
	saml "github.com/MacroPower/provider-authentik/internal/controller/provider/saml"
	scim "github.com/MacroPower/provider-authentik/internal/controller/provider/scim"
	providerconfig "github.com/MacroPower/provider-authentik/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		application.Setup,
		outpost.Setup,
		serviceconnectionkubernetes.Setup,
		blueprint.Setup,
		scopemapping.Setup,
		group.Setup,
		user.Setup,
		flow.Setup,
		stagebinding.Setup,
		ldap.Setup,
		oauth2.Setup,
		proxy.Setup,
		radius.Setup,
		saml.Setup,
		scim.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
