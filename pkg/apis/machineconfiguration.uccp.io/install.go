package machineconfiguration

import (
	machineconfigurationv1 "github.com/uccps-samples/machine-config-operator/pkg/apis/machineconfiguration.uccp.io/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// GroupName defines the API group for machineconfiguration.
const GroupName = "machineconfiguration.uccp.io"

var (
	SchemeBuilder = runtime.NewSchemeBuilder(machineconfigurationv1.Install)
	// Install is a function which adds every version of this group to a scheme
	Install = SchemeBuilder.AddToScheme
)
