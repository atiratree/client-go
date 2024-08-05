// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/openshift/api/authorization/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testing "k8s.io/client-go/testing"
)

// FakeLocalSubjectAccessReviews implements LocalSubjectAccessReviewInterface
type FakeLocalSubjectAccessReviews struct {
	Fake *FakeAuthorizationV1
	ns   string
}

var localsubjectaccessreviewsResource = v1.SchemeGroupVersion.WithResource("localsubjectaccessreviews")

var localsubjectaccessreviewsKind = v1.SchemeGroupVersion.WithKind("LocalSubjectAccessReview")

// Create takes the representation of a localSubjectAccessReview and creates it.  Returns the server's representation of the subjectAccessReviewResponse, and an error, if there is any.
func (c *FakeLocalSubjectAccessReviews) Create(ctx context.Context, localSubjectAccessReview *v1.LocalSubjectAccessReview, opts metav1.CreateOptions) (result *v1.SubjectAccessReviewResponse, err error) {
	emptyResult := &v1.SubjectAccessReviewResponse{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(localsubjectaccessreviewsResource, c.ns, localSubjectAccessReview, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.SubjectAccessReviewResponse), err
}
