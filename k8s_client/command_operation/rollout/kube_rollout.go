package rollout

import (
	"fmt"
	"k8s_CICD/k8s_client/command_operation/rollout/deployment"
	"k8s_CICD/k8s_client/config"
	"k8s_CICD/model/kube_param"
	"time"
)

func KubeRollout(param *kube_param.RolloutParameter) error {
	client := config.KubeConfig()
	timeout := time.Duration(int64(param.Timeout)) * time.Minute

	switch param.State {
	case config.Restart:
		err := deployment.RestartDeployment(client, param.Namespace, param.ResourceName)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		return nil
	case config.Status:
		err := deployment.CheckRolloutStatus(client, param.Namespace, param.ResourceName, timeout)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		return nil
	case config.Undo:
		err := deployment.UndoRollout(client, param.Namespace, param.ResourceName, param.Revision)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		return nil
	default:
		return fmt.Errorf("【暂不支持该操作】: rollout %d", param.State)
	}
}
