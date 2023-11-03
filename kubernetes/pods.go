package kubernetes

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pod struct {
	Name     string
	Status   string
	Creation metav1.Time
}

type Deployment struct {
	Name     string
	Creation string
}

type DaemonSet struct {
	Name     string
	Creation string
}

type ReplicaSet struct {
	Name     string
	Creation string
}

func (opts *KubeConfigOptions) GetPods() ([]Pod, error) {
	var err error
	var podList []Pod

	pods, err := opts.Client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, pod := range pods.Items {
		podList = append(podList, Pod{
			Name:     pod.Name,
			Status:   string(pod.Status.Phase),
			Creation: pod.CreationTimestamp,
		})
	}

	return podList, nil
}

func (opts *KubeConfigOptions) GetDeployments() ([]Deployment, error) {
	var err error
	var deployList []Deployment

	deployments, err := opts.Client.AppsV1().Deployments("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, deployment := range deployments.Items {
		deployList = append(deployList, Deployment{Name: deployment.Name, Creation: deployment.CreationTimestamp.GoString()})
	}

	return deployList, nil
}

func (opts *KubeConfigOptions) GetDaemonSets() ([]DaemonSet, error) {
	var err error
	var daemonsetList []DaemonSet

	daemonsets, err := opts.Client.AppsV1().DaemonSets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, daemonset := range daemonsets.Items {
		daemonsetList = append(daemonsetList, DaemonSet{Name: daemonset.Name, Creation: daemonset.CreationTimestamp.GoString()})
	}

	return daemonsetList, nil
}

func (opts *KubeConfigOptions) GetReplicaSets() ([]ReplicaSet, error) {
	var err error
	var replicasetList []ReplicaSet

	daemonsets, err := opts.Client.AppsV1().ReplicaSets("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, daemonset := range daemonsets.Items {
		replicasetList = append(replicasetList, ReplicaSet{Name: daemonset.Name, Creation: daemonset.CreationTimestamp.GoString()})
	}

	return replicasetList, nil
}
