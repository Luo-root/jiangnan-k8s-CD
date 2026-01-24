package kube_param

import (
	"k8s_CICD/k8s_client/config"
)

type KubeCommendParameter struct {
	OperationType config.Operation `json:"OperationType"`
	Namespace     string           `json:"namespace"`
}

func (p KubeCommendParameter) GetOperationType() config.Operation {
	return p.OperationType
}
