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

type StackObservation struct {

	// Name of the Alertmanager instance configured for this stack.
	AlertmanagerName *string `json:"alertmanagerName,omitempty" tf:"alertmanager_name,omitempty"`

	// Status of the Alertmanager instance configured for this stack.
	AlertmanagerStatus *string `json:"alertmanagerStatus,omitempty" tf:"alertmanager_status,omitempty"`

	// Base URL of the Alertmanager instance configured for this stack.
	AlertmanagerURL *string `json:"alertmanagerUrl,omitempty" tf:"alertmanager_url,omitempty"`

	// User ID of the Alertmanager instance configured for this stack.
	AlertmanagerUserID *float64 `json:"alertmanagerUserId,omitempty" tf:"alertmanager_user_id,omitempty"`

	GraphiteName *string `json:"graphiteName,omitempty" tf:"graphite_name,omitempty"`

	GraphiteStatus *string `json:"graphiteStatus,omitempty" tf:"graphite_status,omitempty"`

	GraphiteURL *string `json:"graphiteUrl,omitempty" tf:"graphite_url,omitempty"`

	GraphiteUserID *float64 `json:"graphiteUserId,omitempty" tf:"graphite_user_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	LogsName *string `json:"logsName,omitempty" tf:"logs_name,omitempty"`

	LogsStatus *string `json:"logsStatus,omitempty" tf:"logs_status,omitempty"`

	LogsURL *string `json:"logsUrl,omitempty" tf:"logs_url,omitempty"`

	LogsUserID *float64 `json:"logsUserId,omitempty" tf:"logs_user_id,omitempty"`

	// Organization id to assign to this stack.
	OrgID *float64 `json:"orgId,omitempty" tf:"org_id,omitempty"`

	// Organization name to assign to this stack.
	OrgName *string `json:"orgName,omitempty" tf:"org_name,omitempty"`

	// Organization slug to assign to this stack.
	OrgSlug *string `json:"orgSlug,omitempty" tf:"org_slug,omitempty"`

	// Prometheus name for this instance.
	PrometheusName *string `json:"prometheusName,omitempty" tf:"prometheus_name,omitempty"`

	// Use this URL to query hosted metrics data e.g. Prometheus data source in Grafana
	PrometheusRemoteEndpoint *string `json:"prometheusRemoteEndpoint,omitempty" tf:"prometheus_remote_endpoint,omitempty"`

	// Use this URL to send prometheus metrics to Grafana cloud
	PrometheusRemoteWriteEndpoint *string `json:"prometheusRemoteWriteEndpoint,omitempty" tf:"prometheus_remote_write_endpoint,omitempty"`

	// Prometheus status for this instance.
	PrometheusStatus *string `json:"prometheusStatus,omitempty" tf:"prometheus_status,omitempty"`

	// Prometheus url for this instance.
	PrometheusURL *string `json:"prometheusUrl,omitempty" tf:"prometheus_url,omitempty"`

	// Prometheus user ID. Used for e.g. remote_write.
	PrometheusUserID *float64 `json:"prometheusUserId,omitempty" tf:"prometheus_user_id,omitempty"`

	// Status of the stack.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TracesName *string `json:"tracesName,omitempty" tf:"traces_name,omitempty"`

	TracesStatus *string `json:"tracesStatus,omitempty" tf:"traces_status,omitempty"`

	TracesURL *string `json:"tracesUrl,omitempty" tf:"traces_url,omitempty"`

	TracesUserID *float64 `json:"tracesUserId,omitempty" tf:"traces_user_id,omitempty"`
}

type StackParameters struct {

	// Description of stack.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Name of stack. Conventionally matches the url of the instance (e.g. “<stack_slug>.grafana.net”).
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region slug to assign to this stack. Changing region will destroy the existing stack and create a new one in the desired region. Use the region list API to get the list of available regions: https://grafana.com/docs/grafana-cloud/reference/cloud-api/#list-regions.
	// +kubebuilder:validation:Optional
	RegionSlug *string `json:"regionSlug,omitempty" tf:"region_slug,omitempty"`

	// Subdomain that the Grafana instance will be available at (i.e. setting slug to “<stack_slug>” will make the instance
	// available at “https://<stack_slug>.grafana.net".
	// +kubebuilder:validation:Required
	Slug *string `json:"slug" tf:"slug,omitempty"`

	// Custom URL for the Grafana instance. Must have a CNAME setup to point to `.grafana.net` before creating the stack
	// +kubebuilder:validation:Optional
	URL *string `json:"url,omitempty" tf:"url,omitempty"`

	// Whether to wait for readiness of the stack after creating it. The check is a HEAD request to the stack URL (Grafana instance). Defaults to `true`.
	// +kubebuilder:validation:Optional
	WaitForReadiness *bool `json:"waitForReadiness,omitempty" tf:"wait_for_readiness,omitempty"`

	// How long to wait for readiness (if enabled). Defaults to `5m0s`.
	// +kubebuilder:validation:Optional
	WaitForReadinessTimeout *string `json:"waitForReadinessTimeout,omitempty" tf:"wait_for_readiness_timeout,omitempty"`
}

// StackSpec defines the desired state of Stack
type StackSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackParameters `json:"forProvider"`
}

// StackStatus defines the observed state of Stack.
type StackStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Stack is the Schema for the Stacks API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackSpec   `json:"spec"`
	Status            StackStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackList contains a list of Stacks
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Stack `json:"items"`
}

// Repository type metadata.
var (
	Stack_Kind             = "Stack"
	Stack_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Stack_Kind}.String()
	Stack_KindAPIVersion   = Stack_Kind + "." + CRDGroupVersion.String()
	Stack_GroupVersionKind = CRDGroupVersion.WithKind(Stack_Kind)
)

func init() {
	SchemeBuilder.Register(&Stack{}, &StackList{})
}
