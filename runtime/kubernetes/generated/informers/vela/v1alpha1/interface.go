// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/go-vela/worker/runtime/kubernetes/generated/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// PipelinePodsTemplates returns a PipelinePodsTemplateInformer.
	PipelinePodsTemplates() PipelinePodsTemplateInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// PipelinePodsTemplates returns a PipelinePodsTemplateInformer.
func (v *version) PipelinePodsTemplates() PipelinePodsTemplateInformer {
	return &pipelinePodsTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
