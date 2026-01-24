package deployment

import (
	"context"
	"encoding/json"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
)

// UpdateDeploymentImage 实现 kubectl apply 修改 Deployment 镜像的功能（Patch 方式）
// 参数：
//   - client: K8s 客户端
//   - namespace: Deployment 命名空间
//   - deployName: Deployment 名称
//   - containerName: 要修改的容器名（Deployment 中的容器名）
//   - newImage: 新镜像地址（如 nginx:1.25.3）
func UpdateDeploymentImage(client *kubernetes.Clientset, namespace, deployName, containerName, newImage string) error {
	// 1. 构造 Patch 数据（只修改指定容器的镜像）
	patchPayload := map[string]interface{}{
		"spec": map[string]interface{}{
			"template": map[string]interface{}{
				"spec": map[string]interface{}{
					"containers": []map[string]interface{}{
						{
							"name":  containerName,
							"image": newImage,
						},
					},
				},
			},
		},
	}
	// 转换为 JSON 字节流
	patchData, err := json.Marshal(patchPayload)
	if err != nil {
		return fmt.Errorf("构造 Patch 数据失败: %v", err)
	}

	// 2. 执行 StrategicMergePatch（和 kubectl apply 底层一致）
	_, err = client.AppsV1().Deployments(namespace).Patch(
		context.TODO(),
		deployName,
		types.StrategicMergePatchType, // 声明式更新核心 Patch 类型
		patchData,
		metav1.PatchOptions{
			FieldManager: "go-client-apply", // 标识更新者，符合 K8s 最佳实践
		},
	)
	if err != nil {
		return fmt.Errorf("更新 Deployment 镜像失败: %v", err)
	}

	fmt.Printf("✅ Deployment %s/%s 的容器 %s 镜像已更新为: %s\n", namespace, deployName, containerName, newImage)
	return nil
}
