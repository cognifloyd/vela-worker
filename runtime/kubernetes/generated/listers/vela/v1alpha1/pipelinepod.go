// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/go-vela/worker/runtime/kubernetes/apis/vela/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PipelinePodLister helps list PipelinePods.
// All objects returned here must be treated as read-only.
type PipelinePodLister interface {
	// List lists all PipelinePods in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PipelinePod, err error)
	// PipelinePods returns an object that can list and get PipelinePods.
	PipelinePods(namespace string) PipelinePodNamespaceLister
	PipelinePodListerExpansion
}

// pipelinePodLister implements the PipelinePodLister interface.
type pipelinePodLister struct {
	indexer cache.Indexer
}

// NewPipelinePodLister returns a new PipelinePodLister.
func NewPipelinePodLister(indexer cache.Indexer) PipelinePodLister {
	return &pipelinePodLister{indexer: indexer}
}

// List lists all PipelinePods in the indexer.
func (s *pipelinePodLister) List(selector labels.Selector) (ret []*v1alpha1.PipelinePod, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PipelinePod))
	})
	return ret, err
}

// PipelinePods returns an object that can list and get PipelinePods.
func (s *pipelinePodLister) PipelinePods(namespace string) PipelinePodNamespaceLister {
	return pipelinePodNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PipelinePodNamespaceLister helps list and get PipelinePods.
// All objects returned here must be treated as read-only.
type PipelinePodNamespaceLister interface {
	// List lists all PipelinePods in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.PipelinePod, err error)
	// Get retrieves the PipelinePod from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.PipelinePod, error)
	PipelinePodNamespaceListerExpansion
}

// pipelinePodNamespaceLister implements the PipelinePodNamespaceLister
// interface.
type pipelinePodNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PipelinePods in the indexer for a given namespace.
func (s pipelinePodNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PipelinePod, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PipelinePod))
	})
	return ret, err
}

// Get retrieves the PipelinePod from the indexer for a given namespace and name.
func (s pipelinePodNamespaceLister) Get(name string) (*v1alpha1.PipelinePod, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("pipelinepod"), name)
	}
	return obj.(*v1alpha1.PipelinePod), nil
}
