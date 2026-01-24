package config

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func KubeConfig() *kubernetes.Clientset {
	// 1. 加载 K8s 配置（优先使用 Pod 内的内置配置，本地开发时用 ~/.kube/config）
	config, err := rest.InClusterConfig()
	if err != nil {
		panic("【kubeconfig配置加载失败】:" + err.Error())
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic("【创建客户端失败】:" + err.Error())
	}
	return client
}
