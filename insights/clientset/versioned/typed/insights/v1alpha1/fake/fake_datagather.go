// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	v1alpha1 "github.com/openshift/api/insights/v1alpha1"
	insightsv1alpha1 "github.com/openshift/client-go/insights/applyconfigurations/insights/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDataGathers implements DataGatherInterface
type FakeDataGathers struct {
	Fake *FakeInsightsV1alpha1
}

var datagathersResource = v1alpha1.SchemeGroupVersion.WithResource("datagathers")

var datagathersKind = v1alpha1.SchemeGroupVersion.WithKind("DataGather")

// Get takes name of the dataGather, and returns the corresponding dataGather object, and an error if there is any.
func (c *FakeDataGathers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DataGather, err error) {
	emptyResult := &v1alpha1.DataGather{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(datagathersResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DataGather), err
}

// List takes label and field selectors, and returns the list of DataGathers that match those selectors.
func (c *FakeDataGathers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DataGatherList, err error) {
	emptyResult := &v1alpha1.DataGatherList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(datagathersResource, datagathersKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DataGatherList{ListMeta: obj.(*v1alpha1.DataGatherList).ListMeta}
	for _, item := range obj.(*v1alpha1.DataGatherList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dataGathers.
func (c *FakeDataGathers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(datagathersResource, opts))
}

// Create takes the representation of a dataGather and creates it.  Returns the server's representation of the dataGather, and an error, if there is any.
func (c *FakeDataGathers) Create(ctx context.Context, dataGather *v1alpha1.DataGather, opts v1.CreateOptions) (result *v1alpha1.DataGather, err error) {
	emptyResult := &v1alpha1.DataGather{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(datagathersResource, dataGather, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DataGather), err
}

// Update takes the representation of a dataGather and updates it. Returns the server's representation of the dataGather, and an error, if there is any.
func (c *FakeDataGathers) Update(ctx context.Context, dataGather *v1alpha1.DataGather, opts v1.UpdateOptions) (result *v1alpha1.DataGather, err error) {
	emptyResult := &v1alpha1.DataGather{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(datagathersResource, dataGather, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DataGather), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDataGathers) UpdateStatus(ctx context.Context, dataGather *v1alpha1.DataGather, opts v1.UpdateOptions) (result *v1alpha1.DataGather, err error) {
	emptyResult := &v1alpha1.DataGather{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(datagathersResource, "status", dataGather, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DataGather), err
}

// Delete takes name of the dataGather and deletes it. Returns an error if one occurs.
func (c *FakeDataGathers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(datagathersResource, name, opts), &v1alpha1.DataGather{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDataGathers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(datagathersResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DataGatherList{})
	return err
}

// Patch applies the patch and returns the patched dataGather.
func (c *FakeDataGathers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DataGather, err error) {
	emptyResult := &v1alpha1.DataGather{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(datagathersResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DataGather), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied dataGather.
func (c *FakeDataGathers) Apply(ctx context.Context, dataGather *insightsv1alpha1.DataGatherApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DataGather, err error) {
	if dataGather == nil {
		return nil, fmt.Errorf("dataGather provided to Apply must not be nil")
	}
	data, err := json.Marshal(dataGather)
	if err != nil {
		return nil, err
	}
	name := dataGather.Name
	if name == nil {
		return nil, fmt.Errorf("dataGather.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.DataGather{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(datagathersResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions()), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DataGather), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeDataGathers) ApplyStatus(ctx context.Context, dataGather *insightsv1alpha1.DataGatherApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.DataGather, err error) {
	if dataGather == nil {
		return nil, fmt.Errorf("dataGather provided to Apply must not be nil")
	}
	data, err := json.Marshal(dataGather)
	if err != nil {
		return nil, err
	}
	name := dataGather.Name
	if name == nil {
		return nil, fmt.Errorf("dataGather.Name must be provided to Apply")
	}
	emptyResult := &v1alpha1.DataGather{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(datagathersResource, *name, types.ApplyPatchType, data, opts.ToPatchOptions(), "status"), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.DataGather), err
}
