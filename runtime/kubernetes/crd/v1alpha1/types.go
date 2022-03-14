// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelinePod defines the config for a given worker.
type PipelinePod struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec defines the PipelinePod configuration for Vela Workers.
	// +optional
	Spec PipelinePodSpec `json:"spec,omitempty"`
}

// PipelinePodSpec configures creation of Pipeline Pods by Vela Workers.
type PipelinePodSpec struct {
	// Template defines defaults for Pipeline Pod creation in Vela Workers.
	Template PipelinePodTemplate `json:"template"`
}

// PipelinePodTemplate describes the data defaults to use when creating each pipeline pod.
type PipelinePodTemplate struct {
	// Meta contains a subset of the standard object metadata (see: metav1.ObjectMeta).
	// +optional
	Meta PipelinePodTemplateMeta `json:"metadata,omitempty"`

	// Spec contains a subset of the pod configuration options (see: v1.PodSpec).
	// +optional
	Spec PipelinePodTemplateSpec `json:"spec,omitempty"`
}

// PipelinePodTemplateMeta is a subset of metav1.ObjectMeta with meta defaults for pipeline pods.
type PipelinePodTemplateMeta struct {
	// Labels is a key value map of strings to organize and categorize pods.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations is a key value map of strings to store additional info on pods.
	// More info: http://kubernetes.io/docs/user-guide/annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
}

// PipelinePodTemplateSpec is (loosely) a subset of v1.PodSpec with spec defaults for pipeline pods.
type PipelinePodTemplateSpec struct {
	// NodeSelector is a selector which must be true for the pipeline pod to fit on a node.
	// Selector which must match a node's labels for the pod to be scheduled on that node.
	// More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	// +optional
	// +mapType=atomic
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
	// Affinity specifies the pipeline pod's scheduling constraints, if any.
	// +optional
	Affinity *v1.Affinity `json:"affinity,omitempty"`
	// Affinity specifies the pipeline pod's tolerations, if any.
	// +optional
	Tolerations []v1.Toleration `json:"tolerations,omitempty"`

	// DNSPolicy sets DNS policy for the pipeline pod.
	// Defaults to "ClusterFirst".
	// Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.
	// +optional
	DNSPolicy v1.DNSPolicy `json:"dnsPolicy,omitempty"`
	// DNSConfig specifies the DNS parameters of a pod.
	// Parameters specified here will be merged to the generated DNS
	// configuration based on DNSPolicy.
	// +optional
	DNSConfig *v1.PodDNSConfig `json:"dnsConfig,omitempty"`

	// Container defines a limited set of defaults to apply to each PipelinePod container.
	// This is analogous to one entry in v1.PodSpec.Containers.
	Container PipelineContainer `json:"container"`

	// SecurityContext holds pod-level security attributes and common container settings.
	// Optional: Defaults to empty.  See type description for default values of each field.
	// +optional
	SecurityContext *PipelinePodSecurityContext `json:"securityContext,omitempty"`
}

// PipelineContainer has defaults for containers in a PipelinePod.
type PipelineContainer struct {
	// SecurityContext defines the security options the container should be run with.
	// If set, the fields of SecurityContext override the equivalent fields of PodSecurityContext.
	// More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	// +optional
	SecurityContext *PipelineContainerSecurityContext `json:"securityContext,omitempty"`
}

// PipelinePodSecurityContext holds pod-level security attributes and common container settings.
type PipelinePodSecurityContext struct {
	// RunAsNonRoot indicates that the container must run as a non-root user.
	// If true, the Kubelet will validate the image at runtime to ensure that it
	// does not run as UID 0 (root) and fail to start the container if it does.
	// If unset or false, no such validation will be performed.
	// +optional
	RunAsNonRoot *bool `json:"runAsNonRoot,omitempty"`
	// Sysctls hold a list of namespaced sysctls used for the pod. Pods with unsupported
	// sysctls (by the container runtime) might fail to launch.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	Sysctls []v1.Sysctl `json:"sysctls,omitempty"`
}

// PipelineContainerSecurityContext holds container-level security configuration.
type PipelineContainerSecurityContext struct {
	// Capabilities contains the capabilities to add/drop when running containers.
	// Defaults to the default set of capabilities granted by the container runtime.
	// Note that this field cannot be set when spec.os.name is windows.
	// +optional
	Capabilities *v1.Capabilities `json:"capabilities,omitempty"`
}
