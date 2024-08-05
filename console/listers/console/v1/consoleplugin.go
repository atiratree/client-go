// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/console/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ConsolePluginLister helps list ConsolePlugins.
// All objects returned here must be treated as read-only.
type ConsolePluginLister interface {
	// List lists all ConsolePlugins in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ConsolePlugin, err error)
	// Get retrieves the ConsolePlugin from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ConsolePlugin, error)
	ConsolePluginListerExpansion
}

// consolePluginLister implements the ConsolePluginLister interface.
type consolePluginLister struct {
	listers.ResourceIndexer[*v1.ConsolePlugin]
}

// NewConsolePluginLister returns a new ConsolePluginLister.
func NewConsolePluginLister(indexer cache.Indexer) ConsolePluginLister {
	return &consolePluginLister{listers.New[*v1.ConsolePlugin](indexer, v1.Resource("consoleplugin"))}
}
