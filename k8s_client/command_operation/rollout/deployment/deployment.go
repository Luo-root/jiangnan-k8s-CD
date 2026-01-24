package deployment

import (
	"context"
	"errors"
	"fmt"
	"k8s_CICD/k8s_client/command_operation/get/deployment"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
)

// RestartDeployment 实现 rollout restart deployment 的功能
// 参数：client - K8s 客户端；namespace - 命名空间；name - Deployment 名称
func RestartDeployment(client *kubernetes.Clientset, namespace, name string) error {
	// 1. 先获取目标 Deployment
	deploy, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("获取 Deployment 失败: %v", err)
	}

	// 2. 生成重启注解（用时间戳确保每次值不同）
	restartAnnotation := "kubectl.kubernetes.io/restartedAt"
	now := time.Now().Format(time.RFC3339)

	// 3. 初始化注解（防止 nil 指针）
	if deploy.Spec.Template.ObjectMeta.Annotations == nil {
		deploy.Spec.Template.ObjectMeta.Annotations = make(map[string]string)
	}
	deploy.Spec.Template.ObjectMeta.Annotations[restartAnnotation] = now

	// 4. 执行 Patch 更新（只修改注解字段，高效）
	_, err = client.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		return fmt.Errorf("重启 Deployment 失败: %v", err)
	}

	fmt.Printf("Deployment %s/%s 已触发重启\n", namespace, name)
	return nil
}

// CheckRolloutStatus 新版实现（无弃用 API，基于 Context 控制超时）
// 参数：
//   - client: K8s 客户端
//   - namespace: Deployment 命名空间
//   - name: Deployment 名称
//   - timeout: 总超时时间（如 5*time.Minute）
func CheckRolloutStatus(client *kubernetes.Clientset, namespace, name string, timeout time.Duration) error {
	//   - interval: 轮询间隔（推荐 2*time.Second）
	interval := 2 * time.Second
	fmt.Printf("正在监控 Deployment %s/%s 的滚动更新状态（超时时间：%v）...\n", namespace, name, timeout)

	// 1. 创建带超时的上下文（替代弃用的 ErrWaitTimeout）
	ctx, cancel := context.WithTimeout(context.TODO(), timeout)
	defer cancel() // 确保函数退出时释放上下文资源

	// 2. 使用新版 PollUntilContextCancel 替代 PollImmediate
	// 核心逻辑：
	// - 立即执行第一次检查（immediate: true）
	// - 直到上下文超时/取消，或返回 true（更新完成）
	// - 轮询间隔由 interval 控制
	err := wait.PollUntilContextCancel(
		ctx,
		interval,
		true, // 立即执行第一次检查（等价于 PollImmediate）
		func(ctx context.Context) (bool, error) {
			// 传递上下文，支持中途取消
			deploy, err := client.AppsV1().Deployments(namespace).Get(ctx, name, metav1.GetOptions{})
			if err != nil {
				// 遇到错误时：返回 false + 错误，会终止轮询
				return false, fmt.Errorf("获取 Deployment 状态失败: %v", err)
			}

			// 滚动更新完成的核心判断条件（官方标准）
			desiredReplicas := deploy.Status.Replicas
			if desiredReplicas == 0 {
				fmt.Printf("Deployment %s/%s 期望副本数为0，滚动更新完成\n", namespace, name)
				return true, nil
			}

			// 补充 ReadyReplicas 校验（更严格的就绪态判断）
			completed := deploy.Status.UpdatedReplicas == desiredReplicas &&
				deploy.Status.AvailableReplicas == desiredReplicas &&
				deploy.Status.ReadyReplicas == desiredReplicas && // 新增：就绪副本数匹配
				deploy.Status.UnavailableReplicas == 0

			if completed {
				fmt.Printf("✅ Deployment %s/%s 滚动更新完成！\n", namespace, name)
				fmt.Printf("  - 期望副本数: %d\n", desiredReplicas)
				fmt.Printf("  - 可用副本数: %d\n", deploy.Status.AvailableReplicas)
				fmt.Printf("  - 最新版本副本数: %d\n", deploy.Status.UpdatedReplicas)
				fmt.Printf("  - 就绪副本数: %d\n", deploy.Status.ReadyReplicas)
				fmt.Printf("  - 不可用副本数: %d\n", deploy.Status.UnavailableReplicas)
				return true, nil
			}

			// 输出中间进度（带时间戳，便于排查）
			fmt.Printf("[%s] 🔄 进度：更新中 - 已更新 %d/%d 副本，可用 %d/%d 副本，不可用 %d 副本\n",
				time.Now().Format("2006-01-02 15:04:05"),
				deploy.Status.UpdatedReplicas, desiredReplicas,
				deploy.Status.AvailableReplicas, desiredReplicas,
				deploy.Status.UnavailableReplicas)
			return false, nil
		},
	)

	// 3. 处理超时/错误（替代弃用的 ErrWaitTimeout）
	if err != nil {
		// 判断是否是上下文超时错误
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return fmt.Errorf("滚动更新超时（%v），Deployment 可能更新失败", timeout)
		}
		// 其他业务错误
		return fmt.Errorf("滚动更新监控失败: %v", err)
	}

	return nil
}

// UndoRollout 新版回滚实现（无废弃类型）
// revision: 0 表示回滚到上一个版本，>0 表示指定版本
func UndoRollout(client *kubernetes.Clientset, namespace, name string, revision int64) error {
	// 1. 获取当前 Deployment
	_, err := client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("获取 Deployment 失败: %v", err)
	}

	// 2. 如果未指定版本（revision=0），自动找上一个版本
	targetRevision := revision
	if targetRevision == 0 {
		revisions, err := deployment.GetDeploymentRevisions(client, namespace, name)
		if err != nil {
			return fmt.Errorf("获取历史版本失败: %v", err)
		}
		if len(revisions) < 2 {
			return fmt.Errorf("无可用的历史版本可回滚")
		}
		// 取倒数第二个版本（上一个版本）
		targetRevision = revisions[len(revisions)-2]
		fmt.Printf("自动选择回滚版本: %d\n", targetRevision)
	}

	// 3. 构建 Patch 数据（核心：设置 revision 注解触发回滚）
	// 这是 kubectl rollout undo 底层的实现方式
	patchData := []byte(fmt.Sprintf(`{
		"spec": {
			"template": {
				"metadata": {
					"annotations": {
						"deployment.kubernetes.io/revision": "%d"
					}
				}
			}
		}
	}`, targetRevision))

	// 4. 执行 Patch 更新（StrategicMergePatch 是 K8s 推荐的 Patch 类型）
	_, err = client.AppsV1().Deployments(namespace).Patch(
		context.TODO(),
		name,
		types.StrategicMergePatchType,
		patchData,
		metav1.PatchOptions{},
	)
	if err != nil {
		return fmt.Errorf("回滚失败: %v", err)
	}

	fmt.Printf("Deployment %s/%s 已回滚到版本 %d\n", namespace, name, targetRevision)
	return nil
}
