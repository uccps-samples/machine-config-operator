package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	// Enable sha256 in container image references
	_ "crypto/sha256"

	"github.com/golang/glog"
	daemon "github.com/uccps-samples/machine-config-operator/pkg/daemon"
	errors "github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var fromEtcPullSpec bool

const (
	// etcPivotFile is used for 4.1 bootimages and is how the MCD
	// currently communicated with this service.
	etcPivotFile = "/etc/pivot/image-pullspec"
)

var pivotCmd = &cobra.Command{
	Use:                   "pivot",
	DisableFlagsInUseLine: true,
	Short:                 "Allows moving from one OSTree deployment to another",
	Args:                  cobra.MaximumNArgs(1),
	Run:                   Execute,
}

// init executes upon import
func init() {
	rootCmd.AddCommand(pivotCmd)
	pivotCmd.PersistentFlags().BoolVarP(&fromEtcPullSpec, "from-etc-pullspec", "P", false, "Parse /etc/pivot/image-pullspec")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
}

func run(_ *cobra.Command, args []string) (retErr error) {
	flag.Set("logtostderr", "true")
	flag.Parse()

	var container string
	if fromEtcPullSpec || len(args) == 0 {
		fromEtcPullSpec = true
		data, err := ioutil.ReadFile(etcPivotFile)
		if err != nil {
			if os.IsNotExist(err) {
				return fmt.Errorf("No container specified")
			}
			return errors.Wrapf(err, "failed to read from %s", etcPivotFile)
		}
		container = strings.TrimSpace(string(data))
	} else {
		container = args[0]
	}

	client := daemon.NewNodeUpdaterClient()

	osImageContentDir, err := daemon.ExtractOSImage(container)
	if err != nil {
		return err
	}
	changed, err := client.Rebase(container, osImageContentDir)
	if err != nil {
		return err
	}

	// Delete the file now that we successfully rebased
	if fromEtcPullSpec {
		if err := os.Remove(etcPivotFile); err != nil {
			if !os.IsNotExist(err) {
				return errors.Wrapf(err, "failed to delete %s", etcPivotFile)
			}
		}
	}

	if !changed {
		glog.Info("No changes; already at target oscontainer")
	}

	return nil
}

// Execute runs the command
func Execute(cmd *cobra.Command, args []string) {
	err := run(cmd, args)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}
