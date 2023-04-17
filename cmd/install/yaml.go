package install

import (
	k8s "kubeadm-helper/cmd/k8sconfig"
	"os"

	"github.com/ghodss/yaml"
)

func readYaml(yamlName string) *k8s.Config {
	// 读取 YAML 文件内容
	content, err := os.ReadFile(yamlName)
	if err != nil {
		panic(err)
	}

	// 解析 YAML 文件到 Config 结构
	var config *k8s.Config
	err = yaml.Unmarshal(content, &config)
	if err != nil {
		panic(err)
	}

	return config
}
