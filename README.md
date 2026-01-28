# jiangnan-k8s-CD
通过apply和rollout接口在k8s中实现滚动升级
## StatusCode
Success : 0
ParameterFail : 1
ApplyFail : 2
RolloutFail : 3
## Headers
- Authorization: "<key>"
## resourceType:
- Deployment : "Deployment"
- StatefulSet : "StatefulSet"
- DaemonSet : "DaemonSet"
- Namespace : "Namespace"
- Service : "Service"
- Ingress : "Ingress"
## rollout
Timeout: 的单位是分钟
### State :
- Restart : 0
- Status : 1
- Undo : 2
```json
{
  "namespace": "memnetai",
  "resourceType": "Deployment",
  "resourceName": "docs",
  "state": 0,
  "timeout": 10,
  "revision": null
}
```
## apply
```json
{
  "namespace": "memnetai",
  "resourceType": "Deployment",
  "resourceName": "docs",
  "containerName": "docs",
  "image": "<image>"
}
```
### 构建dockerfile


# jiangnanCD

基于 Kubernetes 的持续部署工具，通过 API 接口实现 Deployment 的滚动升级和管理。

## 功能特性

- **Apply 接口**：更新 Deployment 镜像
- **Rollout 接口**：管理 Deployment 滚动更新（重启、状态检查、回滚）
- **状态监控**：实时监控滚动更新进度
- **版本回滚**：支持回滚到指定版本

## 快速开始

### 构建镜像
```bash
docker build -t docker.donglizhiyuan.com/library/k8s-cd:v1 .
```

## API 接口说明

### Apply 接口

更新 Deployment 容器镜像。

**请求地址**: `POST /apply`

**请求参数**:
```json
{
  "namespace": "default",
  "resourceType": "Deployment",
  "resourceName": "spring-app",
  "containerName": "spring-app",
  "image": "<image>"
}
```

**参数说明**:
- `namespace`: 命名空间
- `resourceType`: 资源类型（支持 Deployment）
- `resourceName`: 资源名称
- `containerName`: 容器名称
- `image`: 新镜像地址

### Rollout 接口

管理 Deployment 滚动更新操作。

**请求地址**: `POST /rollout`

**请求参数**:
```json
{
  "namespace": "default",
  "resourceType": "Deployment",
  "resourceName": "spring-app",
  "state": 0,
  "timeout": 10,
  "revision": null
}
```

**参数说明**:
- `namespace`: 命名空间
- `resourceType`: 资源类型（支持 Deployment）
- `resourceName`: 资源名称
- `state`: 操作类型
    - `0` - Restart: 重启 Deployment
    - `1` - Status: 检查滚动更新状态
    - `2` - Undo: 回滚操作
- `timeout`: 超时时间（分钟）
- `revision`: 版本号（回滚时使用）

## 支持的资源类型

| 资源类型 | 常量值 |
|---------|--------|
| Deployment | `"Deployment"` |
| StatefulSet | `"StatefulSet"` |
| DaemonSet | `"DaemonSet"` |
| Namespace | `"Namespace"` |
| Service | `"Service"` |
| Ingress | `"Ingress"` |

## 响应状态码

| 状态码 | 常量 | 说明 |
|-------|------|-----|
| 0 | Success | 操作成功 |
| 1 | ParameterFail | 参数错误 |
| 2 | ApplyFail | Apply 操作失败 |
| 3 | RolloutFail | Rollout 操作失败 |

## 操作示例

### 更新镜像
```bash
curl --location --request POST 'http://localhost:8080/apply' \
--header 'Authorization: jQlvvNLjSSPOjdS0iRoFijfMSjDeA7VE' \
--header 'Content-Type: application/json' \
--data-raw '{
  "namespace": "default",
  "resourceType": "Deployment",
  "resourceName": "spring-app",
  "containerName": "spring-app",
  "image": "<image>"
}'
```
### 检查滚动更新状态
```bash
curl -X POST http://localhost:8080/rollout?Authorization=<key> -H "Content-Type: application/json" -d '{ "namespace": "default", "resourceType": "Deployment", "resourceName": "my-app", "state": 1, "timeout": 5, "revision": 0 }'
```
### 重启 Deployment
```bash
curl --location --request POST 'http://localhost:8080/rollout' \
--header 'Authorization: SwYXmHl5ZInmRmYHcl0W0nCNkZJoTb0u' \
--header 'Content-Type: application/json' \
--data-raw '{
  "namespace": "memnetai",
  "resourceType": "Deployment",
  "resourceName": "docs",
  "state": 0,
  "timeout": 10,
  "revision": null
}'
```
## 技术栈

- Go 1.25+
- Gin Web 框架
- Kubernetes Client-go
- Docker

## 贡献

欢迎提交 Issue 和 Pull Request。

## License

MIT