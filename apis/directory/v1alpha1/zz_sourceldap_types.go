/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type SourceLDAPInitParameters struct {

	// (String) Defaults to “.
	// Defaults to “.
	AdditionalGroupDn *string `json:"additionalGroupDn,omitempty" tf:"additional_group_dn,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	AdditionalUserDn *string `json:"additionalUserDn,omitempty" tf:"additional_user_dn,omitempty"`

	// (String)
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	// (String)
	BindCn *string `json:"bindCn,omitempty" tf:"bind_cn,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Defaults to member.
	// Defaults to `member`.
	GroupMembershipField *string `json:"groupMembershipField,omitempty" tf:"group_membership_field,omitempty"`

	// (String) Defaults to (objectClass=group).
	// Defaults to `(objectClass=group)`.
	GroupObjectFilter *string `json:"groupObjectFilter,omitempty" tf:"group_object_filter,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Defaults to objectSid.
	// Defaults to `objectSid`.
	ObjectUniquenessField *string `json:"objectUniquenessField,omitempty" tf:"object_uniqueness_field,omitempty"`

	// (List of String)
	PropertyMappings []*string `json:"propertyMappings,omitempty" tf:"property_mappings,omitempty"`

	// (List of String)
	PropertyMappingsGroup []*string `json:"propertyMappingsGroup,omitempty" tf:"property_mappings_group,omitempty"`

	// (String)
	ServerURI *string `json:"serverUri,omitempty" tf:"server_uri,omitempty"`

	// (String)
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	StartTLS *bool `json:"startTls,omitempty" tf:"start_tls,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	SyncGroups *bool `json:"syncGroups,omitempty" tf:"sync_groups,omitempty"`

	// (String)
	SyncParentGroup *string `json:"syncParentGroup,omitempty" tf:"sync_parent_group,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	SyncUsers *bool `json:"syncUsers,omitempty" tf:"sync_users,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	SyncUsersPassword *bool `json:"syncUsersPassword,omitempty" tf:"sync_users_password,omitempty"`

	// (String) Generated.
	// Generated.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// (String) Defaults to (objectClass=person).
	// Defaults to `(objectClass=person)`.
	UserObjectFilter *string `json:"userObjectFilter,omitempty" tf:"user_object_filter,omitempty"`

	// (String) Defaults to goauthentik.io/sources/%(slug)s.
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate *string `json:"userPathTemplate,omitempty" tf:"user_path_template,omitempty"`
}

type SourceLDAPObservation struct {

	// (String) Defaults to “.
	// Defaults to “.
	AdditionalGroupDn *string `json:"additionalGroupDn,omitempty" tf:"additional_group_dn,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	AdditionalUserDn *string `json:"additionalUserDn,omitempty" tf:"additional_user_dn,omitempty"`

	// (String)
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	// (String)
	BindCn *string `json:"bindCn,omitempty" tf:"bind_cn,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Defaults to member.
	// Defaults to `member`.
	GroupMembershipField *string `json:"groupMembershipField,omitempty" tf:"group_membership_field,omitempty"`

	// (String) Defaults to (objectClass=group).
	// Defaults to `(objectClass=group)`.
	GroupObjectFilter *string `json:"groupObjectFilter,omitempty" tf:"group_object_filter,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Defaults to objectSid.
	// Defaults to `objectSid`.
	ObjectUniquenessField *string `json:"objectUniquenessField,omitempty" tf:"object_uniqueness_field,omitempty"`

	// (List of String)
	PropertyMappings []*string `json:"propertyMappings,omitempty" tf:"property_mappings,omitempty"`

	// (List of String)
	PropertyMappingsGroup []*string `json:"propertyMappingsGroup,omitempty" tf:"property_mappings_group,omitempty"`

	// (String)
	ServerURI *string `json:"serverUri,omitempty" tf:"server_uri,omitempty"`

	// (String)
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	StartTLS *bool `json:"startTls,omitempty" tf:"start_tls,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	SyncGroups *bool `json:"syncGroups,omitempty" tf:"sync_groups,omitempty"`

	// (String)
	SyncParentGroup *string `json:"syncParentGroup,omitempty" tf:"sync_parent_group,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	SyncUsers *bool `json:"syncUsers,omitempty" tf:"sync_users,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	SyncUsersPassword *bool `json:"syncUsersPassword,omitempty" tf:"sync_users_password,omitempty"`

	// (String) Generated.
	// Generated.
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// (String) Defaults to (objectClass=person).
	// Defaults to `(objectClass=person)`.
	UserObjectFilter *string `json:"userObjectFilter,omitempty" tf:"user_object_filter,omitempty"`

	// (String) Defaults to goauthentik.io/sources/%(slug)s.
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	UserPathTemplate *string `json:"userPathTemplate,omitempty" tf:"user_path_template,omitempty"`
}

type SourceLDAPParameters struct {

	// (String) Defaults to “.
	// Defaults to “.
	// +kubebuilder:validation:Optional
	AdditionalGroupDn *string `json:"additionalGroupDn,omitempty" tf:"additional_group_dn,omitempty"`

	// (String) Defaults to “.
	// Defaults to “.
	// +kubebuilder:validation:Optional
	AdditionalUserDn *string `json:"additionalUserDn,omitempty" tf:"additional_user_dn,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	BaseDn *string `json:"baseDn,omitempty" tf:"base_dn,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	BindCn *string `json:"bindCn,omitempty" tf:"bind_cn,omitempty"`

	// (String, Sensitive)
	// +kubebuilder:validation:Optional
	BindPasswordSecretRef v1.SecretKeySelector `json:"bindPasswordSecretRef" tf:"-"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Defaults to member.
	// Defaults to `member`.
	// +kubebuilder:validation:Optional
	GroupMembershipField *string `json:"groupMembershipField,omitempty" tf:"group_membership_field,omitempty"`

	// (String) Defaults to (objectClass=group).
	// Defaults to `(objectClass=group)`.
	// +kubebuilder:validation:Optional
	GroupObjectFilter *string `json:"groupObjectFilter,omitempty" tf:"group_object_filter,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) Defaults to objectSid.
	// Defaults to `objectSid`.
	// +kubebuilder:validation:Optional
	ObjectUniquenessField *string `json:"objectUniquenessField,omitempty" tf:"object_uniqueness_field,omitempty"`

	// (List of String)
	// +kubebuilder:validation:Optional
	PropertyMappings []*string `json:"propertyMappings,omitempty" tf:"property_mappings,omitempty"`

	// (List of String)
	// +kubebuilder:validation:Optional
	PropertyMappingsGroup []*string `json:"propertyMappingsGroup,omitempty" tf:"property_mappings_group,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	ServerURI *string `json:"serverUri,omitempty" tf:"server_uri,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	StartTLS *bool `json:"startTls,omitempty" tf:"start_tls,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	SyncGroups *bool `json:"syncGroups,omitempty" tf:"sync_groups,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	SyncParentGroup *string `json:"syncParentGroup,omitempty" tf:"sync_parent_group,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	SyncUsers *bool `json:"syncUsers,omitempty" tf:"sync_users,omitempty"`

	// (Boolean) Defaults to true.
	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	SyncUsersPassword *bool `json:"syncUsersPassword,omitempty" tf:"sync_users_password,omitempty"`

	// (String) Generated.
	// Generated.
	// +kubebuilder:validation:Optional
	UUID *string `json:"uuid,omitempty" tf:"uuid,omitempty"`

	// (String) Defaults to (objectClass=person).
	// Defaults to `(objectClass=person)`.
	// +kubebuilder:validation:Optional
	UserObjectFilter *string `json:"userObjectFilter,omitempty" tf:"user_object_filter,omitempty"`

	// (String) Defaults to goauthentik.io/sources/%(slug)s.
	// Defaults to `goauthentik.io/sources/%(slug)s`.
	// +kubebuilder:validation:Optional
	UserPathTemplate *string `json:"userPathTemplate,omitempty" tf:"user_path_template,omitempty"`
}

// SourceLDAPSpec defines the desired state of SourceLDAP
type SourceLDAPSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SourceLDAPParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider SourceLDAPInitParameters `json:"initProvider,omitempty"`
}

// SourceLDAPStatus defines the observed state of SourceLDAP.
type SourceLDAPStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SourceLDAPObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SourceLDAP is the Schema for the SourceLDAPs API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type SourceLDAP struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.baseDn) || (has(self.initProvider) && has(self.initProvider.baseDn))",message="spec.forProvider.baseDn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bindCn) || (has(self.initProvider) && has(self.initProvider.bindCn))",message="spec.forProvider.bindCn is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.bindPasswordSecretRef)",message="spec.forProvider.bindPasswordSecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.serverUri) || (has(self.initProvider) && has(self.initProvider.serverUri))",message="spec.forProvider.serverUri is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.slug) || (has(self.initProvider) && has(self.initProvider.slug))",message="spec.forProvider.slug is a required parameter"
	Spec   SourceLDAPSpec   `json:"spec"`
	Status SourceLDAPStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SourceLDAPList contains a list of SourceLDAPs
type SourceLDAPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SourceLDAP `json:"items"`
}

// Repository type metadata.
var (
	SourceLDAP_Kind             = "SourceLDAP"
	SourceLDAP_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SourceLDAP_Kind}.String()
	SourceLDAP_KindAPIVersion   = SourceLDAP_Kind + "." + CRDGroupVersion.String()
	SourceLDAP_GroupVersionKind = CRDGroupVersion.WithKind(SourceLDAP_Kind)
)

func init() {
	SchemeBuilder.Register(&SourceLDAP{}, &SourceLDAPList{})
}
