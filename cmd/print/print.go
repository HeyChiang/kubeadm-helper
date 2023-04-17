package print

import (
	"fmt"
	k8s "kubeadm-helper/cmd/k8sconfig"
	"os"

	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "print",
	Short: "Print the YAML configuration example",
	Run: func(cmd *cobra.Command, args []string) {
		printYamlExample()
	},
}

func printYamlExample() {
	config := k8s.Config{
		K8s: k8s.K8sConfig{
			KubernetesVersion: "1.26.3",
			ServiceCIDR:       "172.17.1.0/24",
			PodNetworkCIDR:    "172.17.2.0/24",
		},
		Machine: []k8s.MachineInfo{
			{Name: "master", Address: "172.1.1.1", Port: 22, User: "root", Password: "password1"},
			{Name: "node1", Address: "172.1.1.1", Port: 22, User: "root", Password: "password1"},
			{Name: "node2", Address: "172.1.1.2", Port: 22, User: "root", Password: "password2"},
		},
		Group: k8s.GroupConfig{
			Etcd:         []string{"master", "node1", "node2"},
			ControlPlane: []string{"master", "node1", "node2"},
			Worker:       []string{"node1", "node2"},
		},
	}

	yamlData, err := yaml.Marshal(config)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("output.yaml", yamlData, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("YAML file created as 'output.yaml'")
}
