package service

import (
	"k8s_CICD/k8s_client/command_operation/apply"
	"k8s_CICD/model/kube_param"
)

func ApplyService(param *kube_param.ApplyParameter) error {
	err := apply.KubeApply(param)
	return err
}
