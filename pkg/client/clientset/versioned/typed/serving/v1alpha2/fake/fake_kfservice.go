/*
Copyright 2019 kubeflow.org.

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
	v1alpha2 "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKFServices implements KFServiceInterface
type FakeKFServices struct {
	Fake *FakeServingV1alpha2
	ns   string
}

var kfservicesResource = schema.GroupVersionResource{Group: "serving.kubeflow.org", Version: "v1alpha2", Resource: "kfservices"}

var kfservicesKind = schema.GroupVersionKind{Group: "serving.kubeflow.org", Version: "v1alpha2", Kind: "KFService"}

// Get takes name of the kFService, and returns the corresponding kFService object, and an error if there is any.
func (c *FakeKFServices) Get(name string, options v1.GetOptions) (result *v1alpha2.KFService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kfservicesResource, c.ns, name), &v1alpha2.KFService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.KFService), err
}

// List takes label and field selectors, and returns the list of KFServices that match those selectors.
func (c *FakeKFServices) List(opts v1.ListOptions) (result *v1alpha2.KFServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kfservicesResource, kfservicesKind, c.ns, opts), &v1alpha2.KFServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.KFServiceList{ListMeta: obj.(*v1alpha2.KFServiceList).ListMeta}
	for _, item := range obj.(*v1alpha2.KFServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kFServices.
func (c *FakeKFServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kfservicesResource, c.ns, opts))

}

// Create takes the representation of a kFService and creates it.  Returns the server's representation of the kFService, and an error, if there is any.
func (c *FakeKFServices) Create(kFService *v1alpha2.KFService) (result *v1alpha2.KFService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kfservicesResource, c.ns, kFService), &v1alpha2.KFService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.KFService), err
}

// Update takes the representation of a kFService and updates it. Returns the server's representation of the kFService, and an error, if there is any.
func (c *FakeKFServices) Update(kFService *v1alpha2.KFService) (result *v1alpha2.KFService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kfservicesResource, c.ns, kFService), &v1alpha2.KFService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.KFService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKFServices) UpdateStatus(kFService *v1alpha2.KFService) (*v1alpha2.KFService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kfservicesResource, "status", c.ns, kFService), &v1alpha2.KFService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.KFService), err
}

// Delete takes name of the kFService and deletes it. Returns an error if one occurs.
func (c *FakeKFServices) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(kfservicesResource, c.ns, name), &v1alpha2.KFService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKFServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kfservicesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.KFServiceList{})
	return err
}

// Patch applies the patch and returns the patched kFService.
func (c *FakeKFServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.KFService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kfservicesResource, c.ns, name, data, subresources...), &v1alpha2.KFService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.KFService), err
}
