/*
Copyright 2020 The Kubernetes Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	k8scnicncfiov1 "github.com/s1061123/k8s-crd-sample/pkg/apis/k8s.cni.cncf.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSampleCRDs implements SampleCRDInterface
type FakeSampleCRDs struct {
	Fake *FakeK8sCniCncfIoV1
	ns   string
}

var samplecrdsResource = schema.GroupVersionResource{Group: "k8s.cni.cncf.io", Version: "v1", Resource: "sample-crds"}

var samplecrdsKind = schema.GroupVersionKind{Group: "k8s.cni.cncf.io", Version: "v1", Kind: "SampleCRD"}

// Get takes name of the sampleCRD, and returns the corresponding sampleCRD object, and an error if there is any.
func (c *FakeSampleCRDs) Get(ctx context.Context, name string, options v1.GetOptions) (result *k8scnicncfiov1.SampleCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(samplecrdsResource, c.ns, name), &k8scnicncfiov1.SampleCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8scnicncfiov1.SampleCRD), err
}

// List takes label and field selectors, and returns the list of SampleCRDs that match those selectors.
func (c *FakeSampleCRDs) List(ctx context.Context, opts v1.ListOptions) (result *k8scnicncfiov1.SampleCRDList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(samplecrdsResource, samplecrdsKind, c.ns, opts), &k8scnicncfiov1.SampleCRDList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &k8scnicncfiov1.SampleCRDList{ListMeta: obj.(*k8scnicncfiov1.SampleCRDList).ListMeta}
	for _, item := range obj.(*k8scnicncfiov1.SampleCRDList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sampleCRDs.
func (c *FakeSampleCRDs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(samplecrdsResource, c.ns, opts))

}

// Create takes the representation of a sampleCRD and creates it.  Returns the server's representation of the sampleCRD, and an error, if there is any.
func (c *FakeSampleCRDs) Create(ctx context.Context, sampleCRD *k8scnicncfiov1.SampleCRD, opts v1.CreateOptions) (result *k8scnicncfiov1.SampleCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(samplecrdsResource, c.ns, sampleCRD), &k8scnicncfiov1.SampleCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8scnicncfiov1.SampleCRD), err
}

// Update takes the representation of a sampleCRD and updates it. Returns the server's representation of the sampleCRD, and an error, if there is any.
func (c *FakeSampleCRDs) Update(ctx context.Context, sampleCRD *k8scnicncfiov1.SampleCRD, opts v1.UpdateOptions) (result *k8scnicncfiov1.SampleCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(samplecrdsResource, c.ns, sampleCRD), &k8scnicncfiov1.SampleCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8scnicncfiov1.SampleCRD), err
}

// Delete takes name of the sampleCRD and deletes it. Returns an error if one occurs.
func (c *FakeSampleCRDs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(samplecrdsResource, c.ns, name), &k8scnicncfiov1.SampleCRD{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSampleCRDs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(samplecrdsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &k8scnicncfiov1.SampleCRDList{})
	return err
}

// Patch applies the patch and returns the patched sampleCRD.
func (c *FakeSampleCRDs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *k8scnicncfiov1.SampleCRD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(samplecrdsResource, c.ns, name, pt, data, subresources...), &k8scnicncfiov1.SampleCRD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*k8scnicncfiov1.SampleCRD), err
}
