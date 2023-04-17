package k8sconfig

// Config 安装k8s的配置信息
type Config struct {
	K8s     K8sConfig     `yaml:"k8s"`
	Machine []MachineInfo `yaml:"machine"`
	Group   GroupConfig   `yaml:"group"`
}

// K8sConfig k8s的配置信息
type K8sConfig struct {
	KubernetesVersion string `yaml:"kubernetes-version"`
	CalicoVersion     string `yaml:"calico-version"`
	ServiceCIDR       string `yaml:"service-cidr"`
	PodNetworkCIDR    string `yaml:"pod-network-cidr"`
}

// MachineInfo k8s的机器配置
type MachineInfo struct {
	Name     string `yaml:"name"`
	Address  string `yaml:"address"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// GroupConfig pod安装的分布
type GroupConfig struct {
	Etcd         []string `yaml:"etcd"`
	ControlPlane []string `yaml:"control-plane"`
	Worker       []string `yaml:"worker"`
}
