/*
Copyright The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.io/sample-controller/pkg/apis/wgpolicyk8s.io/v1alpha1"
)

// FakeClusterPolicyReports implements ClusterPolicyReportInterface
type FakeClusterPolicyReports struct {
	Fake *FakeWgpolicyk8sV1alpha1
}

var clusterpolicyreportsResource = schema.GroupVersionResource{Group: "wgpolicyk8s.io", Version: "v1alpha1", Resource: "clusterpolicyreports"}

var clusterpolicyreportsKind = schema.GroupVersionKind{Group: "wgpolicyk8s.io", Version: "v1alpha1", Kind: "ClusterPolicyReport"}

// Get takes name of the clusterPolicyReport, and returns the corresponding clusterPolicyReport object, and an error if there is any.
func (c *FakeClusterPolicyReports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterPolicyReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterpolicyreportsResource, name), &v1alpha1.ClusterPolicyReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPolicyReport), err
}

// List takes label and field selectors, and returns the list of ClusterPolicyReports that match those selectors.
func (c *FakeClusterPolicyReports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterPolicyReportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterpolicyreportsResource, clusterpolicyreportsKind, opts), &v1alpha1.ClusterPolicyReportList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterPolicyReportList{ListMeta: obj.(*v1alpha1.ClusterPolicyReportList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterPolicyReportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPolicyReports.
func (c *FakeClusterPolicyReports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterpolicyreportsResource, opts))
}

// Create takes the representation of a clusterPolicyReport and creates it.  Returns the server's representation of the clusterPolicyReport, and an error, if there is any.
func (c *FakeClusterPolicyReports) Create(ctx context.Context, clusterPolicyReport *v1alpha1.ClusterPolicyReport, opts v1.CreateOptions) (result *v1alpha1.ClusterPolicyReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterpolicyreportsResource, clusterPolicyReport), &v1alpha1.ClusterPolicyReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPolicyReport), err
}

// Update takes the representation of a clusterPolicyReport and updates it. Returns the server's representation of the clusterPolicyReport, and an error, if there is any.
func (c *FakeClusterPolicyReports) Update(ctx context.Context, clusterPolicyReport *v1alpha1.ClusterPolicyReport, opts v1.UpdateOptions) (result *v1alpha1.ClusterPolicyReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterpolicyreportsResource, clusterPolicyReport), &v1alpha1.ClusterPolicyReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPolicyReport), err
}

// Delete takes name of the clusterPolicyReport and deletes it. Returns an error if one occurs.
func (c *FakeClusterPolicyReports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterpolicyreportsResource, name), &v1alpha1.ClusterPolicyReport{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPolicyReports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterpolicyreportsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterPolicyReportList{})
	return err
}

// Patch applies the patch and returns the patched clusterPolicyReport.
func (c *FakeClusterPolicyReports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterPolicyReport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterpolicyreportsResource, name, pt, data, subresources...), &v1alpha1.ClusterPolicyReport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterPolicyReport), err
}
