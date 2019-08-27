/*
Copyright 2019 The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "sigs.k8s.io/apiserver-builder-alpha/example/podlogs/pkg/apis/podlogs/v1"
)

// PodLogsLister helps list PodLogses.
type PodLogsLister interface {
	// List lists all PodLogses in the indexer.
	List(selector labels.Selector) (ret []*v1.PodLogs, err error)
	// PodLogses returns an object that can list and get PodLogses.
	PodLogses(namespace string) PodLogsNamespaceLister
	PodLogsListerExpansion
}

// podLogsLister implements the PodLogsLister interface.
type podLogsLister struct {
	indexer cache.Indexer
}

// NewPodLogsLister returns a new PodLogsLister.
func NewPodLogsLister(indexer cache.Indexer) PodLogsLister {
	return &podLogsLister{indexer: indexer}
}

// List lists all PodLogses in the indexer.
func (s *podLogsLister) List(selector labels.Selector) (ret []*v1.PodLogs, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PodLogs))
	})
	return ret, err
}

// PodLogses returns an object that can list and get PodLogses.
func (s *podLogsLister) PodLogses(namespace string) PodLogsNamespaceLister {
	return podLogsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodLogsNamespaceLister helps list and get PodLogses.
type PodLogsNamespaceLister interface {
	// List lists all PodLogses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.PodLogs, err error)
	// Get retrieves the PodLogs from the indexer for a given namespace and name.
	Get(name string) (*v1.PodLogs, error)
	PodLogsNamespaceListerExpansion
}

// podLogsNamespaceLister implements the PodLogsNamespaceLister
// interface.
type podLogsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodLogses in the indexer for a given namespace.
func (s podLogsNamespaceLister) List(selector labels.Selector) (ret []*v1.PodLogs, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PodLogs))
	})
	return ret, err
}

// Get retrieves the PodLogs from the indexer for a given namespace and name.
func (s podLogsNamespaceLister) Get(name string) (*v1.PodLogs, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("podlogs"), name)
	}
	return obj.(*v1.PodLogs), nil
}