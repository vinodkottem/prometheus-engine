// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/GoogleCloudPlatform/prometheus-engine/pkg/operator/apis/monitoring/v1"
	scheme "github.com/GoogleCloudPlatform/prometheus-engine/pkg/operator/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeMonitoringsGetter has a method to return a NodeMonitoringInterface.
// A group's client should implement this interface.
type NodeMonitoringsGetter interface {
	NodeMonitorings() NodeMonitoringInterface
}

// NodeMonitoringInterface has methods to work with NodeMonitoring resources.
type NodeMonitoringInterface interface {
	Create(ctx context.Context, nodeMonitoring *v1.NodeMonitoring, opts metav1.CreateOptions) (*v1.NodeMonitoring, error)
	Update(ctx context.Context, nodeMonitoring *v1.NodeMonitoring, opts metav1.UpdateOptions) (*v1.NodeMonitoring, error)
	UpdateStatus(ctx context.Context, nodeMonitoring *v1.NodeMonitoring, opts metav1.UpdateOptions) (*v1.NodeMonitoring, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.NodeMonitoring, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.NodeMonitoringList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NodeMonitoring, err error)
	NodeMonitoringExpansion
}

// nodeMonitorings implements NodeMonitoringInterface
type nodeMonitorings struct {
	client rest.Interface
}

// newNodeMonitorings returns a NodeMonitorings
func newNodeMonitorings(c *MonitoringV1Client) *nodeMonitorings {
	return &nodeMonitorings{
		client: c.RESTClient(),
	}
}

// Get takes name of the nodeMonitoring, and returns the corresponding nodeMonitoring object, and an error if there is any.
func (c *nodeMonitorings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.NodeMonitoring, err error) {
	result = &v1.NodeMonitoring{}
	err = c.client.Get().
		Resource("nodemonitorings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeMonitorings that match those selectors.
func (c *nodeMonitorings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.NodeMonitoringList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.NodeMonitoringList{}
	err = c.client.Get().
		Resource("nodemonitorings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeMonitorings.
func (c *nodeMonitorings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("nodemonitorings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a nodeMonitoring and creates it.  Returns the server's representation of the nodeMonitoring, and an error, if there is any.
func (c *nodeMonitorings) Create(ctx context.Context, nodeMonitoring *v1.NodeMonitoring, opts metav1.CreateOptions) (result *v1.NodeMonitoring, err error) {
	result = &v1.NodeMonitoring{}
	err = c.client.Post().
		Resource("nodemonitorings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeMonitoring).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a nodeMonitoring and updates it. Returns the server's representation of the nodeMonitoring, and an error, if there is any.
func (c *nodeMonitorings) Update(ctx context.Context, nodeMonitoring *v1.NodeMonitoring, opts metav1.UpdateOptions) (result *v1.NodeMonitoring, err error) {
	result = &v1.NodeMonitoring{}
	err = c.client.Put().
		Resource("nodemonitorings").
		Name(nodeMonitoring.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeMonitoring).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *nodeMonitorings) UpdateStatus(ctx context.Context, nodeMonitoring *v1.NodeMonitoring, opts metav1.UpdateOptions) (result *v1.NodeMonitoring, err error) {
	result = &v1.NodeMonitoring{}
	err = c.client.Put().
		Resource("nodemonitorings").
		Name(nodeMonitoring.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(nodeMonitoring).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the nodeMonitoring and deletes it. Returns an error if one occurs.
func (c *nodeMonitorings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("nodemonitorings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeMonitorings) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("nodemonitorings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched nodeMonitoring.
func (c *nodeMonitorings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.NodeMonitoring, err error) {
	result = &v1.NodeMonitoring{}
	err = c.client.Patch(pt).
		Resource("nodemonitorings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
