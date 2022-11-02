// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	operatorv1 "github.com/uccps-samples/api/operator/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSes implements DNSInterface
type FakeDNSes struct {
	Fake *FakeOperatorV1
}

var dnsesResource = schema.GroupVersionResource{Group: "operator.uccp.io", Version: "v1", Resource: "dnses"}

var dnsesKind = schema.GroupVersionKind{Group: "operator.uccp.io", Version: "v1", Kind: "DNS"}

// Get takes name of the dNS, and returns the corresponding dNS object, and an error if there is any.
func (c *FakeDNSes) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.DNS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dnsesResource, name), &operatorv1.DNS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.DNS), err
}

// List takes label and field selectors, and returns the list of DNSes that match those selectors.
func (c *FakeDNSes) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.DNSList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dnsesResource, dnsesKind, opts), &operatorv1.DNSList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.DNSList{ListMeta: obj.(*operatorv1.DNSList).ListMeta}
	for _, item := range obj.(*operatorv1.DNSList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSes.
func (c *FakeDNSes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dnsesResource, opts))
}

// Create takes the representation of a dNS and creates it.  Returns the server's representation of the dNS, and an error, if there is any.
func (c *FakeDNSes) Create(ctx context.Context, dNS *operatorv1.DNS, opts v1.CreateOptions) (result *operatorv1.DNS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dnsesResource, dNS), &operatorv1.DNS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.DNS), err
}

// Update takes the representation of a dNS and updates it. Returns the server's representation of the dNS, and an error, if there is any.
func (c *FakeDNSes) Update(ctx context.Context, dNS *operatorv1.DNS, opts v1.UpdateOptions) (result *operatorv1.DNS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dnsesResource, dNS), &operatorv1.DNS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.DNS), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDNSes) UpdateStatus(ctx context.Context, dNS *operatorv1.DNS, opts v1.UpdateOptions) (*operatorv1.DNS, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dnsesResource, "status", dNS), &operatorv1.DNS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.DNS), err
}

// Delete takes name of the dNS and deletes it. Returns an error if one occurs.
func (c *FakeDNSes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(dnsesResource, name, opts), &operatorv1.DNS{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dnsesResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.DNSList{})
	return err
}

// Patch applies the patch and returns the patched dNS.
func (c *FakeDNSes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.DNS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dnsesResource, name, pt, data, subresources...), &operatorv1.DNS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.DNS), err
}