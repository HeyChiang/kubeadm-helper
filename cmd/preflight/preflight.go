package preflight

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "preflight",
	Short: "Check if the system meets the requirements to install Kubernetes",
	Long:  `This command checks if the system meets the memory, CPU, and other requirements to install Kubernetes.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking system configuration...")
		// Add your code here to preflight the system configuration
	},
}
