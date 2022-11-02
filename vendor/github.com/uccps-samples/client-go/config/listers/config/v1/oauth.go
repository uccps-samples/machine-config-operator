// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/uccps-samples/api/config/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OAuthLister helps list OAuths.
// All objects returned here must be treated as read-only.
type OAuthLister interface {
	// List lists all OAuths in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.OAuth, err error)
	// Get retrieves the OAuth from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.OAuth, error)
	OAuthListerExpansion
}

// oAuthLister implements the OAuthLister interface.
type oAuthLister struct {
	indexer cache.Indexer
}

// NewOAuthLister returns a new OAuthLister.
func NewOAuthLister(indexer cache.Indexer) OAuthLister {
	return &oAuthLister{indexer: indexer}
}

// List lists all OAuths in the indexer.
func (s *oAuthLister) List(selector labels.Selector) (ret []*v1.OAuth, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.OAuth))
	})
	return ret, err
}

// Get retrieves the OAuth from the index for a given name.
func (s *oAuthLister) Get(name string) (*v1.OAuth, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("oauth"), name)
	}
	return obj.(*v1.OAuth), nil
}