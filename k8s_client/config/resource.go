package config

type Resource string

const (
	Deployment  Resource = "Deployment"
	StatefulSet Resource = "StatefulSet"
	DaemonSet   Resource = "DaemonSet"
	Namespace   Resource = "Namespace"
	Service     Resource = "Service"
	Ingress     Resource = "Ingress"
)
