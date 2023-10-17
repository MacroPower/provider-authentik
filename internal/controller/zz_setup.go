/*
Copyright 2021 Upbound Inc.
*/

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/upbound/upjet/pkg/controller"

	application "github.com/MacroPower/provider-authentik/internal/controller/application/application"
	flow "github.com/MacroPower/provider-authentik/internal/controller/flow/flow"
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
		flow.Setup,
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
