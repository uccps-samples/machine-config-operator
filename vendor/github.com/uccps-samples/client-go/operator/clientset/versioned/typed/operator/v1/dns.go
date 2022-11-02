// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/uccps-samples/api/operator/v1"
	scheme "github.com/uccps-samples/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DNSesGetter has a method to return a DNSInterface.
// A group's client should implement this interface.
type DNSesGetter interface {
	DNSes() DNSInterface
}

// DNSInterface has methods to work with DNS resources.
type DNSInterface interface {
	Create(ctx context.Context, dNS *v1.DNS, opts metav1.CreateOptions) (*v1.DNS, error)
	Update(ctx context.Context, dNS *v1.DNS, opts metav1.UpdateOptions) (*v1.DNS, error)
	UpdateStatus(ctx context.Context, dNS *v1.DNS, opts metav1.UpdateOptions) (*v1.DNS, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DNS, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DNSList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DNS, err error)
	DNSExpansion
}

// dNSes implements DNSInterface
type dNSes struct {
	client rest.Interface
}

// newDNSes returns a DNSes
func newDNSes(c *OperatorV1Client) *dNSes {
	return &dNSes{
		client: c.RESTClient(),
	}
}

// Get takes name of the dNS, and returns the corresponding dNS object, and an error if there is any.
func (c *dNSes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DNS, err error) {
	result = &v1.DNS{}
	err = c.client.Get().
		Resource("dnses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DNSes that match those selectors.
func (c *dNSes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DNSList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DNSList{}
	err = c.client.Get().
		Resource("dnses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dNSes.
func (c *dNSes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("dnses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dNS and creates it.  Returns the server's representation of the dNS, and an error, if there is any.
func (c *dNSes) Create(ctx context.Context, dNS *v1.DNS, opts metav1.CreateOptions) (result *v1.DNS, err error) {
	result = &v1.DNS{}
	err = c.client.Post().
		Resource("dnses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNS).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dNS and updates it. Returns the server's representation of the dNS, and an error, if there is any.
func (c *dNSes) Update(ctx context.Context, dNS *v1.DNS, opts metav1.UpdateOptions) (result *v1.DNS, err error) {
	result = &v1.DNS{}
	err = c.client.Put().
		Resource("dnses").
		Name(dNS.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNS).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dNSes) UpdateStatus(ctx context.Context, dNS *v1.DNS, opts metav1.UpdateOptions) (result *v1.DNS, err error) {
	result = &v1.DNS{}
	err = c.client.Put().
		Resource("dnses").
		Name(dNS.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dNS).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dNS and deletes it. Returns an error if one occurs.
func (c *dNSes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("dnses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dNSes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("dnses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dNS.
func (c *dNSes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DNS, err error) {
	result = &v1.DNS{}
	err = c.client.Patch(pt).
		Resource("dnses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
