package service

import (
	"k8s_CICD/k8s_client/command_operation"
	"k8s_CICD/model/kube_param/command_model"
)

func ApplyService(param *command_model.ApplyParameter) error {
	err := command_operation.KubeCommand(param)
	return err
}
