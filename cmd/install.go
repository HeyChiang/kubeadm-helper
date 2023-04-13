package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install kubeadm and other Kubernetes tools",
	Long:  `This command installs kubeadm, kubelet, kubectl and other required Kubernetes tools on your system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing Kubernetes tools...")
		// Add your code here to install Kubernetes tools
	},
}
