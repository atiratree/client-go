// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"k8s.io/client-go/rest"

	v1 "github.com/openshift/api/authorization/v1"
	scheme "github.com/openshift/client-go/authorization/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SelfSubjectRulesReviewsGetter has a method to return a SelfSubjectRulesReviewInterface.
// A group's client should implement this interface.
type SelfSubjectRulesReviewsGetter interface {
	SelfSubjectRulesReviews(namespace string) SelfSubjectRulesReviewInterface
}

// SelfSubjectRulesReviewInterface has methods to work with SelfSubjectRulesReview resources.
type SelfSubjectRulesReviewInterface interface {
	Create(ctx context.Context, selfSubjectRulesReview *v1.SelfSubjectRulesReview, opts metav1.CreateOptions) (*v1.SelfSubjectRulesReview, error)
	SelfSubjectRulesReviewExpansion
}

// selfSubjectRulesReviews implements SelfSubjectRulesReviewInterface
type selfSubjectRulesReviews struct {
	rest.Interface
	namespace string
}

// newSelfSubjectRulesReviews returns a SelfSubjectRulesReviews
func newSelfSubjectRulesReviews(c *AuthorizationV1Client, namespace string) *selfSubjectRulesReviews {
	return &selfSubjectRulesReviews{c.RESTClient(), namespace}
}

// Create takes the representation of a selfSubjectRulesReview and creates it.  Returns the server's representation of the selfSubjectRulesReview, and an error, if there is any.
func (c *selfSubjectRulesReviews) Create(ctx context.Context, selfSubjectRulesReview *v1.SelfSubjectRulesReview, opts metav1.CreateOptions) (result *v1.SelfSubjectRulesReview, err error) {
	result = &v1.SelfSubjectRulesReview{}
	err = c.Interface.Post().
		Namespace(c.namespace).
		Resource("selfsubjectrulesreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selfSubjectRulesReview).
		Do(ctx).
		Into(result)
	return
}
