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

type CertificateKeyPairInitParameters struct {

	// (String)
	CertificateData *string `json:"certificateData,omitempty" tf:"certificate_data,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CertificateKeyPairObservation struct {

	// (String)
	CertificateData *string `json:"certificateData,omitempty" tf:"certificate_data,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String)
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type CertificateKeyPairParameters struct {

	// (String)
	// +kubebuilder:validation:Optional
	CertificateData *string `json:"certificateData,omitempty" tf:"certificate_data,omitempty"`

	// (String, Sensitive)
	// +kubebuilder:validation:Optional
	KeyDataSecretRef *v1.SecretKeySelector `json:"keyDataSecretRef,omitempty" tf:"-"`

	// (String)
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// CertificateKeyPairSpec defines the desired state of CertificateKeyPair
type CertificateKeyPairSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CertificateKeyPairParameters `json:"forProvider"`
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
	InitProvider CertificateKeyPairInitParameters `json:"initProvider,omitempty"`
}

// CertificateKeyPairStatus defines the observed state of CertificateKeyPair.
type CertificateKeyPairStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CertificateKeyPairObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateKeyPair is the Schema for the CertificateKeyPairs API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,authentik}
type CertificateKeyPair struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.certificateData) || (has(self.initProvider) && has(self.initProvider.certificateData))",message="spec.forProvider.certificateData is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   CertificateKeyPairSpec   `json:"spec"`
	Status CertificateKeyPairStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CertificateKeyPairList contains a list of CertificateKeyPairs
type CertificateKeyPairList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateKeyPair `json:"items"`
}

// Repository type metadata.
var (
	CertificateKeyPair_Kind             = "CertificateKeyPair"
	CertificateKeyPair_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CertificateKeyPair_Kind}.String()
	CertificateKeyPair_KindAPIVersion   = CertificateKeyPair_Kind + "." + CRDGroupVersion.String()
	CertificateKeyPair_GroupVersionKind = CRDGroupVersion.WithKind(CertificateKeyPair_Kind)
)

func init() {
	SchemeBuilder.Register(&CertificateKeyPair{}, &CertificateKeyPairList{})
}