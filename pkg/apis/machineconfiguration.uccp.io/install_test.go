package machineconfiguration

import (
	"testing"
)

func TestGroupName(t *testing.T) {
	if got, want := GroupName, "machineconfiguration.uccp.io"; got != want {
		t.Fatalf("mismatch group name, got: %s want: %s", got, want)
	}
}
