package e2e_test

import (
	"context"
	"testing"

	"github.com/uccps-samples/machine-config-operator/test/framework"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestOSImageURL(t *testing.T) {
	cs := framework.NewClientSet("")

	// grab the latest worker- MC
	mcp, err := cs.MachineConfigPools().Get(context.TODO(), "worker", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%#v", err)
	}

	mc, err := cs.MachineConfigs().Get(context.TODO(), mcp.Status.Configuration.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%#v", err)
	}

	if mc.Spec.OSImageURL == "" {
		t.Fatalf("Empty OSImageURL for %s", mc.Name)
	}

	// grab the latest master- MC
	mcp, err = cs.MachineConfigPools().Get(context.TODO(), "master", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%#v", err)
	}
	mc, err = cs.MachineConfigs().Get(context.TODO(), mcp.Status.Configuration.Name, metav1.GetOptions{})
	if err != nil {
		t.Fatalf("%#v", err)
	}

	if mc.Spec.OSImageURL == "" {
		t.Fatalf("Empty OSImageURL for %s", mc.Name)
	}
}
