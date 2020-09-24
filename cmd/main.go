package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"k8s.io/klog/v2"
)

var ()

var rootCmd = &cobra.Command{
	Use:     "subscription-injection-webhook",
	Version: "0.0.1",
	Short:   "",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		klog.Info("Hello World Webhook!")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	klog.InitFlags(nil)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	cobra.OnInitialize()

	rootCmd.Flags().AddGoFlagSet(flag.CommandLine)
}
