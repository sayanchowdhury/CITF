/*
Copyright 2018 The OpenEBS Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/openebs/CITF/pkg/apis/openebs.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StoragePoolLister helps list StoragePools.
type StoragePoolLister interface {
	// List lists all StoragePools in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StoragePool, err error)
	// Get retrieves the StoragePool from the index for a given name.
	Get(name string) (*v1alpha1.StoragePool, error)
	StoragePoolListerExpansion
}

// storagePoolLister implements the StoragePoolLister interface.
type storagePoolLister struct {
	indexer cache.Indexer
}

// NewStoragePoolLister returns a new StoragePoolLister.
func NewStoragePoolLister(indexer cache.Indexer) StoragePoolLister {
	return &storagePoolLister{indexer: indexer}
}

// List lists all StoragePools in the indexer.
func (s *storagePoolLister) List(selector labels.Selector) (ret []*v1alpha1.StoragePool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoragePool))
	})
	return ret, err
}

// Get retrieves the StoragePool from the index for a given name.
func (s *storagePoolLister) Get(name string) (*v1alpha1.StoragePool, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagepool"), name)
	}
	return obj.(*v1alpha1.StoragePool), nil
}
