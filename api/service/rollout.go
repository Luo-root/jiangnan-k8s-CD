package service

import (
	"k8s_CICD/k8s_client/command_operation/rollout"
	"k8s_CICD/model/kube_param"
)

func RolloutService(param *kube_param.RolloutParameter) error {
	err := rollout.KubeRollout(param)
	return err
}
