package kubernetes

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (opts *KubeConfigOptions) GetPodCount() int {
	if opts.Client != nil {
		pods, err := opts.Client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return 0
		}

		return len(pods.Items)
	} else {
		return 0
	}
}

func (opts *KubeConfigOptions) GetPodStatuses() (error, [][]string) {
	var err error
	var data [][]string

	pods, err := opts.Client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err, nil
	}

	for _, pod := range pods.Items {
		data = append(data, []string{fmt.Sprintf("%s", pod.Status.Phase), pod.Name})
	}

	return nil, data
}
