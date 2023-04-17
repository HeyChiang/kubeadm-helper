package install

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "install",
	Short: "Install kubeadm and other Kubernetes print",
	Long:  `This command installs kubeadm, kubelet, kubectl and other required Kubernetes print on your system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing Kubernetes print...")
		// Add your code here to install Kubernetes print
	},
}
