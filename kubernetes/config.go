package kubernetes

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func (opts *KubeConfigOptions) CreateConfig() error {
	var err error

	opts.Config = &rest.Config{
		Host:        "",
		BearerToken: "",
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
	}

	opts.Client, err = kubernetes.NewForConfig(opts.Config)
	if err != nil {
		return err
	}

	return nil
}

func (opts *KubeConfigOptions) GetNamespaces() error {
	var err error

	namespaceList, err := opts.Client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, n := range namespaceList.Items {
		opts.Namespaces = append(opts.Namespaces, n.Name)
	}

	return nil
}
