/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
	v1alpha1 "github.com/MacroPower/provider-authentik/apis/authentik/v1alpha1"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	errors "github.com/pkg/errors"
	resource "github.com/upbound/upjet/pkg/resource"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// ResolveReferences of this Group.
func (mg *Group) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var mrsp reference.MultiResolutionResponse
	var err error

	mrsp, err = r.ResolveMultiple(ctx, reference.MultiResolutionRequest{
		CurrentValues: reference.FromFloatPtrValues(mg.Spec.ForProvider.Users),
		Extract:       resource.ExtractParamPath("id", true),
		References:    mg.Spec.ForProvider.UsersRefs,
		Selector:      mg.Spec.ForProvider.UsersSelector,
		To: reference.To{
			List:    &UserList{},
			Managed: &User{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.Users")
	}
	mg.Spec.ForProvider.Users = reference.ToFloatPtrValues(mrsp.ResolvedValues)
	mg.Spec.ForProvider.UsersRefs = mrsp.ResolvedReferences

	return nil
}

// ResolveReferences of this SourceOAuth.
func (mg *SourceOAuth) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthenticationFlow),
		Extract:      resource.ExtractParamPath("uuid", true),
		Reference:    mg.Spec.ForProvider.AuthenticationFlowRef,
		Selector:     mg.Spec.ForProvider.AuthenticationFlowSelector,
		To: reference.To{
			List:    &v1alpha1.FlowList{},
			Managed: &v1alpha1.Flow{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthenticationFlow")
	}
	mg.Spec.ForProvider.AuthenticationFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthenticationFlowRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnrollmentFlow),
		Extract:      resource.ExtractParamPath("uuid", true),
		Reference:    mg.Spec.ForProvider.EnrollmentFlowRef,
		Selector:     mg.Spec.ForProvider.EnrollmentFlowSelector,
		To: reference.To{
			List:    &v1alpha1.FlowList{},
			Managed: &v1alpha1.Flow{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnrollmentFlow")
	}
	mg.Spec.ForProvider.EnrollmentFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnrollmentFlowRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SourcePlex.
func (mg *SourcePlex) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthenticationFlow),
		Extract:      resource.ExtractParamPath("uuid", true),
		Reference:    mg.Spec.ForProvider.AuthenticationFlowRef,
		Selector:     mg.Spec.ForProvider.AuthenticationFlowSelector,
		To: reference.To{
			List:    &v1alpha1.FlowList{},
			Managed: &v1alpha1.Flow{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthenticationFlow")
	}
	mg.Spec.ForProvider.AuthenticationFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthenticationFlowRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnrollmentFlow),
		Extract:      resource.ExtractParamPath("uuid", true),
		Reference:    mg.Spec.ForProvider.EnrollmentFlowRef,
		Selector:     mg.Spec.ForProvider.EnrollmentFlowSelector,
		To: reference.To{
			List:    &v1alpha1.FlowList{},
			Managed: &v1alpha1.Flow{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnrollmentFlow")
	}
	mg.Spec.ForProvider.EnrollmentFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnrollmentFlowRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SourceSAML.
func (mg *SourceSAML) ResolveReferences(ctx context.Context, c client.Reader) error {
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.AuthenticationFlow),
		Extract:      resource.ExtractParamPath("uuid", true),
		Reference:    mg.Spec.ForProvider.AuthenticationFlowRef,
		Selector:     mg.Spec.ForProvider.AuthenticationFlowSelector,
		To: reference.To{
			List:    &v1alpha1.FlowList{},
			Managed: &v1alpha1.Flow{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.AuthenticationFlow")
	}
	mg.Spec.ForProvider.AuthenticationFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.AuthenticationFlowRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.EnrollmentFlow),
		Extract:      resource.ExtractParamPath("uuid", true),
		Reference:    mg.Spec.ForProvider.EnrollmentFlowRef,
		Selector:     mg.Spec.ForProvider.EnrollmentFlowSelector,
		To: reference.To{
			List:    &v1alpha1.FlowList{},
			Managed: &v1alpha1.Flow{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.EnrollmentFlow")
	}
	mg.Spec.ForProvider.EnrollmentFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.EnrollmentFlowRef = rsp.ResolvedReference

	rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
		CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.PreAuthenticationFlow),
		Extract:      resource.ExtractParamPath("uuid", true),
		Reference:    mg.Spec.ForProvider.PreAuthenticationFlowRef,
		Selector:     mg.Spec.ForProvider.PreAuthenticationFlowSelector,
		To: reference.To{
			List:    &v1alpha1.FlowList{},
			Managed: &v1alpha1.Flow{},
		},
	})
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.PreAuthenticationFlow")
	}
	mg.Spec.ForProvider.PreAuthenticationFlow = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.PreAuthenticationFlowRef = rsp.ResolvedReference

	return nil
}
