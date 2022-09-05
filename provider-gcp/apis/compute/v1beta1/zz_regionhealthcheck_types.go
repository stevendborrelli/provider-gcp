/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RegionHealthCheckGRPCHealthCheckObservation struct {
}

type RegionHealthCheckGRPCHealthCheckParameters struct {

	// The gRPC service name for the health check.
	// The value of grpcServiceName has the following meanings by convention:
	// +kubebuilder:validation:Optional
	GRPCServiceName *string `json:"grpcServiceName,omitempty" tf:"grpc_service_name,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`
}

type RegionHealthCheckHTTPHealthCheckObservation struct {
}

type RegionHealthCheckHTTPHealthCheckParameters struct {

	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckHTTPSHealthCheckObservation struct {
}

type RegionHealthCheckHTTPSHealthCheckParameters struct {

	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckHttp2HealthCheckObservation struct {
}

type RegionHealthCheckHttp2HealthCheckParameters struct {

	// +kubebuilder:validation:Optional
	Host *string `json:"host,omitempty" tf:"host,omitempty"`

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// +kubebuilder:validation:Optional
	RequestPath *string `json:"requestPath,omitempty" tf:"request_path,omitempty"`

	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckLogConfigObservation struct {
}

type RegionHealthCheckLogConfigParameters struct {

	// Indicates whether or not to export logs. This is false by default,
	// which means no health check logging will be done.
	// +kubebuilder:validation:Optional
	Enable *bool `json:"enable,omitempty" tf:"enable,omitempty"`
}

type RegionHealthCheckObservation struct {

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`

	// an identifier for the resource with format projects/{{project}}/regions/{{region}}/healthChecks/{{name}}
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The URI of the created resource.
	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	// The type of the health check. One of HTTP, HTTP2, HTTPS, TCP, or SSL.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type RegionHealthCheckParameters struct {

	// How often  to send a health check. The default value is 5
	// seconds.
	// +kubebuilder:validation:Optional
	CheckIntervalSec *float64 `json:"checkIntervalSec,omitempty" tf:"check_interval_sec,omitempty"`

	// An optional description of this resource. Provide this property when
	// you create the resource.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	GRPCHealthCheck []RegionHealthCheckGRPCHealthCheckParameters `json:"grpcHealthCheck,omitempty" tf:"grpc_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPHealthCheck []RegionHealthCheckHTTPHealthCheckParameters `json:"httpHealthCheck,omitempty" tf:"http_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	HTTPSHealthCheck []RegionHealthCheckHTTPSHealthCheckParameters `json:"httpsHealthCheck,omitempty" tf:"https_health_check,omitempty"`

	// A so-far unhealthy instance will be marked healthy after this many
	// consecutive successes. The default value is 2.
	// +kubebuilder:validation:Optional
	HealthyThreshold *float64 `json:"healthyThreshold,omitempty" tf:"healthy_threshold,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	Http2HealthCheck []RegionHealthCheckHttp2HealthCheckParameters `json:"http2HealthCheck,omitempty" tf:"http2_health_check,omitempty"`

	// Configure logging on this health check.
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	LogConfig []RegionHealthCheckLogConfigParameters `json:"logConfig,omitempty" tf:"log_config,omitempty"`

	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The Region in which the created health check should reside.
	// If it is not provided, the provider region is used.
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"region,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	SSLHealthCheck []RegionHealthCheckSSLHealthCheckParameters `json:"sslHealthCheck,omitempty" tf:"ssl_health_check,omitempty"`

	// A nested object resource
	// Structure is documented below.
	// +kubebuilder:validation:Optional
	TCPHealthCheck []RegionHealthCheckTCPHealthCheckParameters `json:"tcpHealthCheck,omitempty" tf:"tcp_health_check,omitempty"`

	// How long  to wait before claiming failure.
	// The default value is 5 seconds.  It is invalid for timeoutSec to have
	// greater value than checkIntervalSec.
	// +kubebuilder:validation:Optional
	TimeoutSec *float64 `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`

	// A so-far healthy instance will be marked unhealthy after this many
	// consecutive failures. The default value is 2.
	// +kubebuilder:validation:Optional
	UnhealthyThreshold *float64 `json:"unhealthyThreshold,omitempty" tf:"unhealthy_threshold,omitempty"`
}

type RegionHealthCheckSSLHealthCheckObservation struct {
}

type RegionHealthCheckSSLHealthCheckParameters struct {

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// +kubebuilder:validation:Optional
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

type RegionHealthCheckTCPHealthCheckObservation struct {
}

type RegionHealthCheckTCPHealthCheckParameters struct {

	// +kubebuilder:validation:Optional
	Port *float64 `json:"port,omitempty" tf:"port,omitempty"`

	// +kubebuilder:validation:Optional
	PortName *string `json:"portName,omitempty" tf:"port_name,omitempty"`

	// +kubebuilder:validation:Optional
	PortSpecification *string `json:"portSpecification,omitempty" tf:"port_specification,omitempty"`

	// +kubebuilder:validation:Optional
	ProxyHeader *string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`

	// +kubebuilder:validation:Optional
	Request *string `json:"request,omitempty" tf:"request,omitempty"`

	// +kubebuilder:validation:Optional
	Response *string `json:"response,omitempty" tf:"response,omitempty"`
}

// RegionHealthCheckSpec defines the desired state of RegionHealthCheck
type RegionHealthCheckSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RegionHealthCheckParameters `json:"forProvider"`
}

// RegionHealthCheckStatus defines the observed state of RegionHealthCheck.
type RegionHealthCheckStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RegionHealthCheckObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RegionHealthCheck is the Schema for the RegionHealthChecks API. Health Checks determine whether instances are responsive and able to do work.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcp}
type RegionHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RegionHealthCheckSpec   `json:"spec"`
	Status            RegionHealthCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RegionHealthCheckList contains a list of RegionHealthChecks
type RegionHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RegionHealthCheck `json:"items"`
}

// Repository type metadata.
var (
	RegionHealthCheck_Kind             = "RegionHealthCheck"
	RegionHealthCheck_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RegionHealthCheck_Kind}.String()
	RegionHealthCheck_KindAPIVersion   = RegionHealthCheck_Kind + "." + CRDGroupVersion.String()
	RegionHealthCheck_GroupVersionKind = CRDGroupVersion.WithKind(RegionHealthCheck_Kind)
)

func init() {
	SchemeBuilder.Register(&RegionHealthCheck{}, &RegionHealthCheckList{})
}