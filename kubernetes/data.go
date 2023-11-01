package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var kubeConfig = KubeConfigOptions{}

type KubeConfigOptions struct {
	Namespaces []string

	Config *rest.Config
	Client *kubernetes.Clientset
}

func GetKubeConfig() KubeConfigOptions {
	return kubeConfig
}
