// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	machineconfigurationopenshiftiov1 "github.com/uccps-samples/machine-config-operator/pkg/apis/machineconfiguration.uccp.io/v1"
	versioned "github.com/uccps-samples/machine-config-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/uccps-samples/machine-config-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/uccps-samples/machine-config-operator/pkg/generated/listers/machineconfiguration.uccp.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MachineConfigPoolInformer provides access to a shared informer and lister for
// MachineConfigPools.
type MachineConfigPoolInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.MachineConfigPoolLister
}

type machineConfigPoolInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewMachineConfigPoolInformer constructs a new informer for MachineConfigPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineConfigPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMachineConfigPoolInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredMachineConfigPoolInformer constructs a new informer for MachineConfigPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMachineConfigPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineconfigurationV1().MachineConfigPools().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MachineconfigurationV1().MachineConfigPools().Watch(context.TODO(), options)
			},
		},
		&machineconfigurationopenshiftiov1.MachineConfigPool{},
		resyncPeriod,
		indexers,
	)
}

func (f *machineConfigPoolInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMachineConfigPoolInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *machineConfigPoolInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&machineconfigurationopenshiftiov1.MachineConfigPool{}, f.defaultInformer)
}

func (f *machineConfigPoolInformer) Lister() v1.MachineConfigPoolLister {
	return v1.NewMachineConfigPoolLister(f.Informer().GetIndexer())
}
