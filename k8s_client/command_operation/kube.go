package command_operation

import (
	"fmt"
	"k8s_CICD/k8s_client/command/apply"
	"k8s_CICD/k8s_client/command/rollout"
	"k8s_CICD/k8s_client/config"
	"k8s_CICD/model/kube_command_interface"
	"k8s_CICD/model/kube_param/command"
)

func KubeCommand(param kube_command_interface.KubeParam) error {

	switch param.GetOperationType() {
	case config.Apply:
		applyParam, ok := param.(command_model.ApplyParameter)
		if !ok {
			return fmt.Errorf("【类型转换错误不是预期的 ApplyParameter】: command %T", param)
		}
		err := apply.KubeApply(&applyParam)
		if err != nil {
			return err
		}
		return nil
	case config.Rollout:
		rolloutParam, ok := param.(command_model.RolloutParameter)
		if !ok {
			return fmt.Errorf("【类型转换错误不是预期的 RolloutParameter】: command %T", param)
		}
		err := rollout.KubeRollout(&rolloutParam)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("【暂不支持该操作类型】: Command %d", param.GetOperationType())
	}
}
