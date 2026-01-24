package deployment

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetDeploymentRevisions 获取 Deployment 的历史版本列表
// 参数：revision - 可选，指定回滚的版本号（0 表示回滚到上一个版本）
func GetDeploymentRevisions(client *kubernetes.Clientset, namespace, name string) ([]int64, error) {
	// 获取 Deployment 的所有修订版本（通过 ReplicaSet）
	rsList, err := client.AppsV1().ReplicaSets(namespace).List(context.TODO(), metav1.ListOptions{
		LabelSelector: fmt.Sprintf("app.kubernetes.io/name=%s,app.kubernetes.io/instance=%s", name, name), // 根据实际 label 调整
	})
	if err != nil {
		return nil, fmt.Errorf("获取 ReplicaSet 失败: %v", err)
	}

	// 解析每个 RS 的 revision 注解
	var revisions []int64
	for _, rs := range rsList.Items {
		if revStr, ok := rs.Annotations["deployment.kubernetes.io/revision"]; ok {
			rev, err := strconv.ParseInt(revStr, 10, 64)
			if err == nil {
				revisions = append(revisions, rev)
			}
		}
	}

	// 排序版本号
	sort.Slice(revisions, func(i, j int) bool {
		return revisions[i] < revisions[j]
	})
	return revisions, nil
}
