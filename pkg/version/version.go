package version

import (
	"fmt"
)

var (
	// Raw is the string representation of the version. This will be replaced
	// with the calculated version at build time.
	Raw = "v0.0.0-was-not-built-properly"

	// Hash is the git hash we've built the MCO with
	Hash = "was-not-built-properly"

	// String is the human-friendly representation of the version.
	String = fmt.Sprintf("MachineConfigOperator %s", Raw)
)
