package main

import (
	"fmt"
	"kubeadm-helper/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Print("执行命令出错了")
		return
	}
}
