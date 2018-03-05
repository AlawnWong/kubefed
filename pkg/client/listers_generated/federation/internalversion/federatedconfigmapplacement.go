/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package internalversion

import (
	federation "github.com/marun/fnord/pkg/apis/federation"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FederatedConfigMapPlacementLister helps list FederatedConfigMapPlacements.
type FederatedConfigMapPlacementLister interface {
	// List lists all FederatedConfigMapPlacements in the indexer.
	List(selector labels.Selector) (ret []*federation.FederatedConfigMapPlacement, err error)
	// FederatedConfigMapPlacements returns an object that can list and get FederatedConfigMapPlacements.
	FederatedConfigMapPlacements(namespace string) FederatedConfigMapPlacementNamespaceLister
	FederatedConfigMapPlacementListerExpansion
}

// federatedConfigMapPlacementLister implements the FederatedConfigMapPlacementLister interface.
type federatedConfigMapPlacementLister struct {
	indexer cache.Indexer
}

// NewFederatedConfigMapPlacementLister returns a new FederatedConfigMapPlacementLister.
func NewFederatedConfigMapPlacementLister(indexer cache.Indexer) FederatedConfigMapPlacementLister {
	return &federatedConfigMapPlacementLister{indexer: indexer}
}

// List lists all FederatedConfigMapPlacements in the indexer.
func (s *federatedConfigMapPlacementLister) List(selector labels.Selector) (ret []*federation.FederatedConfigMapPlacement, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*federation.FederatedConfigMapPlacement))
	})
	return ret, err
}

// FederatedConfigMapPlacements returns an object that can list and get FederatedConfigMapPlacements.
func (s *federatedConfigMapPlacementLister) FederatedConfigMapPlacements(namespace string) FederatedConfigMapPlacementNamespaceLister {
	return federatedConfigMapPlacementNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FederatedConfigMapPlacementNamespaceLister helps list and get FederatedConfigMapPlacements.
type FederatedConfigMapPlacementNamespaceLister interface {
	// List lists all FederatedConfigMapPlacements in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*federation.FederatedConfigMapPlacement, err error)
	// Get retrieves the FederatedConfigMapPlacement from the indexer for a given namespace and name.
	Get(name string) (*federation.FederatedConfigMapPlacement, error)
	FederatedConfigMapPlacementNamespaceListerExpansion
}

// federatedConfigMapPlacementNamespaceLister implements the FederatedConfigMapPlacementNamespaceLister
// interface.
type federatedConfigMapPlacementNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FederatedConfigMapPlacements in the indexer for a given namespace.
func (s federatedConfigMapPlacementNamespaceLister) List(selector labels.Selector) (ret []*federation.FederatedConfigMapPlacement, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*federation.FederatedConfigMapPlacement))
	})
	return ret, err
}

// Get retrieves the FederatedConfigMapPlacement from the indexer for a given namespace and name.
func (s federatedConfigMapPlacementNamespaceLister) Get(name string) (*federation.FederatedConfigMapPlacement, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(federation.Resource("federatedconfigmapplacement"), name)
	}
	return obj.(*federation.FederatedConfigMapPlacement), nil
}
