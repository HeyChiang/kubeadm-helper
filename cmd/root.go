package cmd

import (
	"github.com/spf13/cobra"
	"kubeadm-helper/cmd/add"
	"kubeadm-helper/cmd/install"
	"kubeadm-helper/cmd/preflight"
	"kubeadm-helper/cmd/print"
)

var rootCmd = &cobra.Command{
	Short: "A tool to install Kubernetes",
	Long:  `This tool helps you install Kubernetes, preflight system configuration and copy certificates.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(preflight.Cmd)
	rootCmd.AddCommand(install.Cmd)
	rootCmd.AddCommand(add.Cmd)
	rootCmd.AddCommand(print.Cmd)
}
