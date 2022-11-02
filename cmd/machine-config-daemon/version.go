package main

import (
	"flag"
	"fmt"

	"github.com/uccps-samples/machine-config-operator/pkg/version"
	"github.com/spf13/cobra"
)

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Machine Config Daemon",
		Long:  `All software has versions. This is Machine Config Daemon's.`,
		Run:   runVersionCmd,
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersionCmd(cmd *cobra.Command, args []string) {
	flag.Set("logtostderr", "true")
	flag.Parse()

	program := "MachineConfigDaemon"
	version := version.Raw + "-" + version.Hash

	fmt.Println(program, version)
}
