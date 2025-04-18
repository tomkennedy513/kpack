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

package fake

import (
	"context"

	v1alpha2 "github.com/pivotal/kpack/pkg/apis/build/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBuilders implements BuilderInterface
type FakeBuilders struct {
	Fake *FakeKpackV1alpha2
	ns   string
}

var buildersResource = v1alpha2.SchemeGroupVersion.WithResource("builders")

var buildersKind = v1alpha2.SchemeGroupVersion.WithKind("Builder")

// Get takes name of the builder, and returns the corresponding builder object, and an error if there is any.
func (c *FakeBuilders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Builder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildersResource, c.ns, name), &v1alpha2.Builder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Builder), err
}

// List takes label and field selectors, and returns the list of Builders that match those selectors.
func (c *FakeBuilders) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.BuilderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildersResource, buildersKind, c.ns, opts), &v1alpha2.BuilderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.BuilderList{ListMeta: obj.(*v1alpha2.BuilderList).ListMeta}
	for _, item := range obj.(*v1alpha2.BuilderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested builders.
func (c *FakeBuilders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildersResource, c.ns, opts))

}

// Create takes the representation of a builder and creates it.  Returns the server's representation of the builder, and an error, if there is any.
func (c *FakeBuilders) Create(ctx context.Context, builder *v1alpha2.Builder, opts v1.CreateOptions) (result *v1alpha2.Builder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildersResource, c.ns, builder), &v1alpha2.Builder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Builder), err
}

// Update takes the representation of a builder and updates it. Returns the server's representation of the builder, and an error, if there is any.
func (c *FakeBuilders) Update(ctx context.Context, builder *v1alpha2.Builder, opts v1.UpdateOptions) (result *v1alpha2.Builder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildersResource, c.ns, builder), &v1alpha2.Builder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Builder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuilders) UpdateStatus(ctx context.Context, builder *v1alpha2.Builder, opts v1.UpdateOptions) (*v1alpha2.Builder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildersResource, "status", c.ns, builder), &v1alpha2.Builder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Builder), err
}

// Delete takes name of the builder and deletes it. Returns an error if one occurs.
func (c *FakeBuilders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(buildersResource, c.ns, name, opts), &v1alpha2.Builder{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuilders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.BuilderList{})
	return err
}

// Patch applies the patch and returns the patched builder.
func (c *FakeBuilders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Builder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildersResource, c.ns, name, pt, data, subresources...), &v1alpha2.Builder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Builder), err
}
