package command_model

import (
	"k8s_CICD/k8s_client/config"
	"k8s_CICD/model/kube_param"
)

type RolloutParameter struct {
	Command      kube_param.KubeCommendParameter
	ResourceType int          `json:"ResourceType"`
	ResourceName string       `json:"ResourceName"`
	State        config.State `json:"state"`
	Timeout      int          `json:"timeout"`
	Revision     int64        `json:"revision"`
}

func (p RolloutParameter) GetOperationType() config.Operation {
	return p.Command.GetOperationType()
}
