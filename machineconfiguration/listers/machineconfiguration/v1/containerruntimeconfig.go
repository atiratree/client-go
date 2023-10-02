// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/machineconfiguration/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ContainerRuntimeConfigLister helps list ContainerRuntimeConfigs.
// All objects returned here must be treated as read-only.
type ContainerRuntimeConfigLister interface {
	// List lists all ContainerRuntimeConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ContainerRuntimeConfig, err error)
	// Get retrieves the ContainerRuntimeConfig from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ContainerRuntimeConfig, error)
	ContainerRuntimeConfigListerExpansion
}

// containerRuntimeConfigLister implements the ContainerRuntimeConfigLister interface.
type containerRuntimeConfigLister struct {
	indexer cache.Indexer
}

// NewContainerRuntimeConfigLister returns a new ContainerRuntimeConfigLister.
func NewContainerRuntimeConfigLister(indexer cache.Indexer) ContainerRuntimeConfigLister {
	return &containerRuntimeConfigLister{indexer: indexer}
}

// List lists all ContainerRuntimeConfigs in the indexer.
func (s *containerRuntimeConfigLister) List(selector labels.Selector) (ret []*v1.ContainerRuntimeConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ContainerRuntimeConfig))
	})
	return ret, err
}

// Get retrieves the ContainerRuntimeConfig from the index for a given name.
func (s *containerRuntimeConfigLister) Get(name string) (*v1.ContainerRuntimeConfig, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("containerruntimeconfig"), name)
	}
	return obj.(*v1.ContainerRuntimeConfig), nil
}
