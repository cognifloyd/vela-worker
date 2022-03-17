// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	velav1alpha1 "github.com/go-vela/worker/runtime/kubernetes/apis/vela/v1alpha1"
	versioned "github.com/go-vela/worker/runtime/kubernetes/generated/clientset/versioned"
	internalinterfaces "github.com/go-vela/worker/runtime/kubernetes/generated/informers/internalinterfaces"
	v1alpha1 "github.com/go-vela/worker/runtime/kubernetes/generated/listers/vela/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PipelinePodsTemplateInformer provides access to a shared informer and lister for
// PipelinePodsTemplates.
type PipelinePodsTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PipelinePodsTemplateLister
}

type pipelinePodsTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPipelinePodsTemplateInformer constructs a new informer for PipelinePodsTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPipelinePodsTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPipelinePodsTemplateInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPipelinePodsTemplateInformer constructs a new informer for PipelinePodsTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPipelinePodsTemplateInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VelaV1alpha1().PipelinePodsTemplates(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VelaV1alpha1().PipelinePodsTemplates(namespace).Watch(context.TODO(), options)
			},
		},
		&velav1alpha1.PipelinePodsTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *pipelinePodsTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPipelinePodsTemplateInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *pipelinePodsTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&velav1alpha1.PipelinePodsTemplate{}, f.defaultInformer)
}

func (f *pipelinePodsTemplateInformer) Lister() v1alpha1.PipelinePodsTemplateLister {
	return v1alpha1.NewPipelinePodsTemplateLister(f.Informer().GetIndexer())
}
