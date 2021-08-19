/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	scheme "github.com/pivotal/kpack/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SourceResolversGetter has a method to return a SourceResolverInterface.
// A group's client should implement this interface.
type SourceResolversGetter interface {
	SourceResolvers(namespace string) SourceResolverInterface
}

// SourceResolverInterface has methods to work with SourceResolver resources.
type SourceResolverInterface interface {
	Create(ctx context.Context, sourceResolver *v1alpha2.SourceResolver, opts v1.CreateOptions) (*v1alpha2.SourceResolver, error)
	Update(ctx context.Context, sourceResolver *v1alpha2.SourceResolver, opts v1.UpdateOptions) (*v1alpha2.SourceResolver, error)
	UpdateStatus(ctx context.Context, sourceResolver *v1alpha2.SourceResolver, opts v1.UpdateOptions) (*v1alpha2.SourceResolver, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.SourceResolver, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.SourceResolverList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.SourceResolver, err error)
	SourceResolverExpansion
}

// sourceResolvers implements SourceResolverInterface
type sourceResolvers struct {
	client rest.Interface
	ns     string
}

// newSourceResolvers returns a SourceResolvers
func newSourceResolvers(c *KpackV1alpha2Client, namespace string) *sourceResolvers {
	return &sourceResolvers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sourceResolver, and returns the corresponding sourceResolver object, and an error if there is any.
func (c *sourceResolvers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.SourceResolver, err error) {
	result = &v1alpha2.SourceResolver{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourceresolvers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SourceResolvers that match those selectors.
func (c *sourceResolvers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.SourceResolverList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.SourceResolverList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sourceresolvers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sourceResolvers.
func (c *sourceResolvers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sourceresolvers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sourceResolver and creates it.  Returns the server's representation of the sourceResolver, and an error, if there is any.
func (c *sourceResolvers) Create(ctx context.Context, sourceResolver *v1alpha2.SourceResolver, opts v1.CreateOptions) (result *v1alpha2.SourceResolver, err error) {
	result = &v1alpha2.SourceResolver{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sourceresolvers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sourceResolver).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sourceResolver and updates it. Returns the server's representation of the sourceResolver, and an error, if there is any.
func (c *sourceResolvers) Update(ctx context.Context, sourceResolver *v1alpha2.SourceResolver, opts v1.UpdateOptions) (result *v1alpha2.SourceResolver, err error) {
	result = &v1alpha2.SourceResolver{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourceresolvers").
		Name(sourceResolver.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sourceResolver).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sourceResolvers) UpdateStatus(ctx context.Context, sourceResolver *v1alpha2.SourceResolver, opts v1.UpdateOptions) (result *v1alpha2.SourceResolver, err error) {
	result = &v1alpha2.SourceResolver{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sourceresolvers").
		Name(sourceResolver.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sourceResolver).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sourceResolver and deletes it. Returns an error if one occurs.
func (c *sourceResolvers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourceresolvers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sourceResolvers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sourceresolvers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sourceResolver.
func (c *sourceResolvers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.SourceResolver, err error) {
	result = &v1alpha2.SourceResolver{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sourceresolvers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}