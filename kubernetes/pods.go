package kubernetes

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pod struct {
	name     string
	status   string
	creation string
}

type Deployment struct {
	name     string
	creation string
}

type DaemonSet struct {
	name     string
	creation string
}

type ReplicaSet struct {
	name     string
	creation string
}

func (opts *KubeConfigOptions) GetPods() ([]Pod, error) {
	var err error
	var podList []Pod

	pods, err := opts.Client.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	for _, pod := range pods.Items {
		podList = append(podList, Pod{name: pod.Name, status: string(pod.Status.Phase), creation: pod.CreationTimestamp.GoString()})
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
		deployList = append(deployList, Deployment{name: deployment.Name, creation: deployment.CreationTimestamp.GoString()})
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
		daemonsetList = append(daemonsetList, DaemonSet{name: daemonset.Name, creation: daemonset.CreationTimestamp.GoString()})
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
		replicasetList = append(replicasetList, ReplicaSet{name: daemonset.Name, creation: daemonset.CreationTimestamp.GoString()})
	}

	return replicasetList, nil
}
