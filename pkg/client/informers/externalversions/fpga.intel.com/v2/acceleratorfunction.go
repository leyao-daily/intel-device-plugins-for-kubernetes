// Copyright 2020 Intel Corporation. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	"context"
	time "time"

	fpgaintelcomv2 "github.com/intel/intel-device-plugins-for-kubernetes/pkg/apis/fpga.intel.com/v2"
	versioned "github.com/intel/intel-device-plugins-for-kubernetes/pkg/client/clientset/versioned"
	internalinterfaces "github.com/intel/intel-device-plugins-for-kubernetes/pkg/client/informers/externalversions/internalinterfaces"
	v2 "github.com/intel/intel-device-plugins-for-kubernetes/pkg/client/listers/fpga.intel.com/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AcceleratorFunctionInformer provides access to a shared informer and lister for
// AcceleratorFunctions.
type AcceleratorFunctionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.AcceleratorFunctionLister
}

type acceleratorFunctionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAcceleratorFunctionInformer constructs a new informer for AcceleratorFunction type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAcceleratorFunctionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAcceleratorFunctionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAcceleratorFunctionInformer constructs a new informer for AcceleratorFunction type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAcceleratorFunctionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FpgaV2().AcceleratorFunctions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FpgaV2().AcceleratorFunctions(namespace).Watch(context.TODO(), options)
			},
		},
		&fpgaintelcomv2.AcceleratorFunction{},
		resyncPeriod,
		indexers,
	)
}

func (f *acceleratorFunctionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAcceleratorFunctionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *acceleratorFunctionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&fpgaintelcomv2.AcceleratorFunction{}, f.defaultInformer)
}

func (f *acceleratorFunctionInformer) Lister() v2.AcceleratorFunctionLister {
	return v2.NewAcceleratorFunctionLister(f.Informer().GetIndexer())
}