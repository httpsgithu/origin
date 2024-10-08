// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ImageContentPolicyLister helps list ImageContentPolicies.
// All objects returned here must be treated as read-only.
type ImageContentPolicyLister interface {
	// List lists all ImageContentPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ImageContentPolicy, err error)
	// Get retrieves the ImageContentPolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ImageContentPolicy, error)
	ImageContentPolicyListerExpansion
}

// imageContentPolicyLister implements the ImageContentPolicyLister interface.
type imageContentPolicyLister struct {
	listers.ResourceIndexer[*v1.ImageContentPolicy]
}

// NewImageContentPolicyLister returns a new ImageContentPolicyLister.
func NewImageContentPolicyLister(indexer cache.Indexer) ImageContentPolicyLister {
	return &imageContentPolicyLister{listers.New[*v1.ImageContentPolicy](indexer, v1.Resource("imagecontentpolicy"))}
}
