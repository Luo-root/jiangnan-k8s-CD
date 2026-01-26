package kube_param

import (
	"k8s_CICD/k8s_client/config"
)

type ApplyParameter struct {
	Namespace     string          `json:"namespace"`
	ResourceType  config.Resource `json:"resourceType"`
	ResourceName  string          `json:"resourceName"`
	ContainerName string          `json:"containerName"`
	Image         string          `json:"image"`
}
