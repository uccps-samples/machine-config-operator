package main

import (
	"flag"
	"os"

	"github.com/uccps-samples/machine-config-operator/pkg/server"
	"github.com/spf13/cobra"
	"k8s.io/component-base/cli"
)

const (
	componentName = "machine-config-server"
)

var (
	rootCmd = &cobra.Command{
		Use:   componentName,
		Short: "Run Machine Config Server",
		Long:  "",
	}

	rootOpts struct {
		sport  int
		isport int
		cert   string
		key    string
	}
)

func init() {
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	rootCmd.PersistentFlags().IntVar(&rootOpts.sport, "secure-port", server.SecurePort, "secure port to serve ignition configs")
	rootCmd.PersistentFlags().StringVar(&rootOpts.cert, "cert", "/etc/ssl/mcs/tls.crt", "cert file for TLS")
	rootCmd.PersistentFlags().StringVar(&rootOpts.key, "key", "/etc/ssl/mcs/tls.key", "key file for TLS")
	rootCmd.PersistentFlags().IntVar(&rootOpts.isport, "insecure-port", server.InsecurePort, "insecure port to serve ignition configs")
}

func main() {
	code := cli.Run(rootCmd)
	os.Exit(code)
}
