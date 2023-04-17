package install

import (
	"fmt"
	"github.com/spf13/cobra"
	k8s "kubeadm-helper/cmd/k8sconfig"
)

var Cmd = &cobra.Command{
	Use:   "install",
	Short: "Install Kubernetes",
	Long:  `This command installs kubeadm, kubelet, kubectl and Kubernetes in your system.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Installing Kubernetes print...")

		if len(args) == 0 {
			print("请")
			return
		}

		// read yaml
		k8sConfig := readYaml(args[0])
		InitNewSSHForMultipleMachines(k8sConfig.Machine)
		containerD()

	},
}

func CheckMachinesInGroups(config k8s.Config) {
	machineMap := make(map[string]k8s.MachineInfo)
	for _, machine := range config.Machine {
		machineMap[machine.Name] = machine
	}

	for group, groupMachines := range map[string][]string{
		"etcd":         config.Group.Etcd,
		"controlPlane": config.Group.ControlPlane,
		"worker":       config.Group.Worker,
	} {
		for _, machineName := range groupMachines {
			if _, ok := machineMap[machineName]; ok {
				fmt.Printf("Machine '%s' matched in group '%s'\n", machineName, group)
			} else {
				fmt.Printf("Machine '%s' not found in machine list\n", machineName)
			}
		}
	}
}
