/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	"context"
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