# jiangnanCD
### resourceType:
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
  "namespace": "default",
  "resourceType": "Deployment",
  "resourceName": "spring-app",
  "state": 0,
  "timeout": 10,
  "revision": null
}
```
## apply
```json
{
  "namespace": "default",
  "resourceType": "Deployment",
  "resourceName": "spring-app",
  "containerName": "spring-app",
  "image": "docker.donglizhiyuan.com/spring-app:v1.0.0"
}
```