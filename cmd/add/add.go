package add

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "add",
	Short: "Copy Kubernetes certificates to different machines",
	Long:  `This command copies Kubernetes master and worker node certificates to different machines in the cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Copying Kubernetes certificates...")
		// Add your code here to copy certificates to different machines
	},
}
