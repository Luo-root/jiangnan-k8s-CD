package command_model

import (
	"k8s_CICD/k8s_client/config"
	"k8s_CICD/model/kube_param"
)

type ApplyParameter struct {
	Command       kube_param.KubeCommendParameter
	ResourceType  config.Resource `json:"ResourceType"`
	ResourceName  string          `json:"ResourceName"`
	ContainerName string          `json:"containerName"`
	Image         string          `json:"image"`
}

func (p ApplyParameter) GetOperationType() config.Operation {
	return p.Command.GetOperationType()
}
