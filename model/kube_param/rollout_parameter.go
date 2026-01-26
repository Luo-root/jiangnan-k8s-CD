package kube_param

import (
	"k8s_CICD/k8s_client/config"
)

type RolloutParameter struct {
	Namespace    string          `json:"namespace"`
	ResourceType config.Resource `json:"resourceType"`
	ResourceName string          `json:"resourceName"`
	State        config.State    `json:"state"`
	Timeout      int             `json:"timeout"`
	Revision     int64           `json:"revision"`
}
