// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	operatorv1 "github.com/uccps-samples/api/operator/v1"
	versioned "github.com/uccps-samples/client-go/operator/clientset/versioned"
	internalinterfaces "github.com/uccps-samples/client-go/operator/informers/externalversions/internalinterfaces"
	v1 "github.com/uccps-samples/client-go/operator/listers/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubeStorageVersionMigratorInformer provides access to a shared informer and lister for
// KubeStorageVersionMigrators.
type KubeStorageVersionMigratorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.KubeStorageVersionMigratorLister
}

type kubeStorageVersionMigratorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewKubeStorageVersionMigratorInformer constructs a new informer for KubeStorageVersionMigrator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeStorageVersionMigratorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeStorageVersionMigratorInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredKubeStorageVersionMigratorInformer constructs a new informer for KubeStorageVersionMigrator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeStorageVersionMigratorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().KubeStorageVersionMigrators().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().KubeStorageVersionMigrators().Watch(context.TODO(), options)
			},
		},
		&operatorv1.KubeStorageVersionMigrator{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeStorageVersionMigratorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeStorageVersionMigratorInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeStorageVersionMigratorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorv1.KubeStorageVersionMigrator{}, f.defaultInformer)
}

func (f *kubeStorageVersionMigratorInformer) Lister() v1.KubeStorageVersionMigratorLister {
	return v1.NewKubeStorageVersionMigratorLister(f.Informer().GetIndexer())
}
