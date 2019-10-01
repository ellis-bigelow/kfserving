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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	servingv1alpha2 "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2"
	versioned "github.com/kubeflow/kfserving/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubeflow/kfserving/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/kubeflow/kfserving/pkg/client/listers/serving/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KFServiceInformer provides access to a shared informer and lister for
// KFServices.
type KFServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.KFServiceLister
}

type kFServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKFServiceInformer constructs a new informer for KFService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKFServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKFServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKFServiceInformer constructs a new informer for KFService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKFServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1alpha2().KFServices(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServingV1alpha2().KFServices(namespace).Watch(options)
			},
		},
		&servingv1alpha2.KFService{},
		resyncPeriod,
		indexers,
	)
}

func (f *kFServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKFServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kFServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servingv1alpha2.KFService{}, f.defaultInformer)
}

func (f *kFServiceInformer) Lister() v1alpha2.KFServiceLister {
	return v1alpha2.NewKFServiceLister(f.Informer().GetIndexer())
}
