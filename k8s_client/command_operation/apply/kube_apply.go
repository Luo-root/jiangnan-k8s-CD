package apply

import (
	"fmt"
	"k8s_CICD/k8s_client/command_operation/apply/deployment"
	"k8s_CICD/k8s_client/config"
	"k8s_CICD/model/kube_param"
)

func KubeApply(param *kube_param.ApplyParameter) error {
	client := config.KubeConfig()
	switch param.ResourceType {
	case config.Deployment:
		err := deployment.UpdateDeploymentImage(client, param.Namespace, param.ResourceName, param.ContainerName, param.Image)
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		return nil
	default:
		return fmt.Errorf(string("【暂不支持该资源类型】: apply " + param.ResourceType))
	}
}
