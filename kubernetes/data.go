package kubernetes

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var kubeConfig = KubeConfigOptions{}

type KubeConfigOptions struct {
	Namespaces []string
	Nodes      []string

	Config *rest.Config
	Client *kubernetes.Clientset
}

func (opts *KubeConfigOptions) CreateConfig() (KubeConfigOptions, error) {
	var err error

	kubeConfig.Config = &rest.Config{
		Host:        "https://10.254.254.11:6443",
		BearerToken: "eyJhbGciOiJSUzI1NiIsImtpZCI6ImVBNFR3SDN5WjZvRGtfNk50UU1XdE5tNzk3SGhfUUJXTEhXT2VaRjBrU2sifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJkZWZhdWx0Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9zZWNyZXQubmFtZSI6InZkZWJvcmdlciIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJ2ZGVib3JnZXIiLCJrdWJlcm5ldGVzLmlvL3NlcnZpY2VhY2NvdW50L3NlcnZpY2UtYWNjb3VudC51aWQiOiJhNjJhZmQ2My0zYWIzLTQ4ODItYjYwOC0wZDk5OTk2MDIzODMiLCJzdWIiOiJzeXN0ZW06c2VydmljZWFjY291bnQ6ZGVmYXVsdDp2ZGVib3JnZXIifQ.ctN9fyY7vVLCo6Z6mflDlb1sww5Uwt2-4SL6blmoixBPYG2Yw5zCclCQy2qm94tzyec3c7hHFe-dHuLf89c2QCjbX90luyKvPaFB18bloInwRDUPIx5QO1Ec25zlYVLaAvZLAIXwv1kkAfed8iDWtbq7OYFf8hfSxQwnJNikFBCthSA6D_rchr8m7X6UwjHXWSN1iZXGjRK8PQuz6VpdEF4_r6jfCqFUc4DRC8l0Xh3WU2_TZu41urg1fV-QEUYJgn8LBka4antBACngmNmpeJ1saJu4CTUKhc9kZZEAfKYEcbWtTJJgT0C9jTjfnQhFdV9NWBmGuCIQXCtyW0wW2w",
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}

	kubeConfig.Client, err = kubernetes.NewForConfig(kubeConfig.Config)
	if err != nil {
		return KubeConfigOptions{}, err
	}

	return kubeConfig, nil
}

func GetKubeConfig() KubeConfigOptions {
	return kubeConfig
}

func UpdateKubeConfig(config KubeConfigOptions) {
	kubeConfig = config
}
