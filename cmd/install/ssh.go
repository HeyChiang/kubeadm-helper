package install

import (
	"fmt"
	k8s "kubeadm-helper/cmd/k8sconfig"
	"log"
	"time"

	"golang.org/x/crypto/ssh"
)

type SSH struct {
	Client *ssh.Client
}

var sessions []*ssh.Session

func NewSSHClient(machine k8s.MachineInfo) (*ssh.Session, error) {
	config := &ssh.ClientConfig{
		User: machine.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(machine.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         5 * time.Second,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", machine.Address, machine.Port), config)
	if err != nil {
		return nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, err
	}

	return session, nil
}

func InitNewSSHForMultipleMachines(machines []k8s.MachineInfo) {

	for _, machine := range machines {
		session, err := NewSSHClient(machine)
		if err != nil {
			panic("connect " + machine.Name + " had a error! " + err.Error())
		}
		sessions = append(sessions, session)
	}
}

func containerD() {
	installCmds := []string{
		"yum-config-manager --add-repo https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo",
		"yum list|grep containerd",
		"yum install -y containerd.io",
		"sudo systemctl enable --now containerd",
	}

	executeCmd(installCmds)
}

func configContainerD() {
	installCmds := []string{
		"containerd config default > /etc/containerd/config.toml",
		"sed -i 's/systemd_cgroup = false/systemd_cgroup = true/g' /etc/containerd/config.toml",
		`echo '[plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
  endpoint = ["https://czerws6q.mirror.aliyuncs.com"]
[plugins."io.containerd.grpc.v1.cri".registry.mirrors."k8s.gcr.io"]
  endpoint = ["https://registry.aliyuncs.com"]
[plugins."io.containerd.grpc.v1.cri".registry.mirrors."registry.k8s.io"]
  endpoint = ["https://k8s.m.daocloud.io"]' >> /etc/containerd/config.toml`,
	}
	executeCmd(installCmds)
}

func executeCmd(commands []string) {
	for _, session := range sessions {
		for _, cmd := range commands {
			err := session.Run(cmd)
			if err != nil {
				log.Printf("Failed to run command '%s': %s", cmd, err)
				continue
			}
			log.Printf("Successfully executed command '%s'", cmd)
		}
	}
}

func closeAllSession(sessions []*ssh.Session) {
	for _, session := range sessions {
		err := session.Close()
		if err != nil {
			log.Printf("close session error！")
			return
		}
	}
}
