package kube_command_interface

import "k8s_CICD/k8s_client/config"

type KubeParam interface {
	GetOperationType() config.Operation
}
