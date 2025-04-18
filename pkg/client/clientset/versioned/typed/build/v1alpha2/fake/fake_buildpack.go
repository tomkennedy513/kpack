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

// FakeBuildpacks implements BuildpackInterface
type FakeBuildpacks struct {
	Fake *FakeKpackV1alpha2
	ns   string
}

var buildpacksResource = v1alpha2.SchemeGroupVersion.WithResource("buildpacks")

var buildpacksKind = v1alpha2.SchemeGroupVersion.WithKind("Buildpack")

// Get takes name of the buildpack, and returns the corresponding buildpack object, and an error if there is any.
func (c *FakeBuildpacks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.Buildpack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(buildpacksResource, c.ns, name), &v1alpha2.Buildpack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Buildpack), err
}

// List takes label and field selectors, and returns the list of Buildpacks that match those selectors.
func (c *FakeBuildpacks) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.BuildpackList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(buildpacksResource, buildpacksKind, c.ns, opts), &v1alpha2.BuildpackList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.BuildpackList{ListMeta: obj.(*v1alpha2.BuildpackList).ListMeta}
	for _, item := range obj.(*v1alpha2.BuildpackList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested buildpacks.
func (c *FakeBuildpacks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(buildpacksResource, c.ns, opts))

}

// Create takes the representation of a buildpack and creates it.  Returns the server's representation of the buildpack, and an error, if there is any.
func (c *FakeBuildpacks) Create(ctx context.Context, buildpack *v1alpha2.Buildpack, opts v1.CreateOptions) (result *v1alpha2.Buildpack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(buildpacksResource, c.ns, buildpack), &v1alpha2.Buildpack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Buildpack), err
}

// Update takes the representation of a buildpack and updates it. Returns the server's representation of the buildpack, and an error, if there is any.
func (c *FakeBuildpacks) Update(ctx context.Context, buildpack *v1alpha2.Buildpack, opts v1.UpdateOptions) (result *v1alpha2.Buildpack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(buildpacksResource, c.ns, buildpack), &v1alpha2.Buildpack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Buildpack), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBuildpacks) UpdateStatus(ctx context.Context, buildpack *v1alpha2.Buildpack, opts v1.UpdateOptions) (*v1alpha2.Buildpack, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(buildpacksResource, "status", c.ns, buildpack), &v1alpha2.Buildpack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Buildpack), err
}

// Delete takes name of the buildpack and deletes it. Returns an error if one occurs.
func (c *FakeBuildpacks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(buildpacksResource, c.ns, name, opts), &v1alpha2.Buildpack{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBuildpacks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(buildpacksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha2.BuildpackList{})
	return err
}

// Patch applies the patch and returns the patched buildpack.
func (c *FakeBuildpacks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.Buildpack, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(buildpacksResource, c.ns, name, pt, data, subresources...), &v1alpha2.Buildpack{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.Buildpack), err
}
