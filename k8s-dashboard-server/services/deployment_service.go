package services

import (
	"context"
	core_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type DeploymentService struct {
	clientset *kubernetes.Clientset
}

func NewDeploymentService(clientset *kubernetes.Clientset) *DeploymentService {
	return &DeploymentService{clientset: clientset}
}

func (ds *DeploymentService) List(namespace string) (podsList *core_v1.PodList, err error) {
	// 获取default命名空间下的所有POD
	return ds.clientset.CoreV1().Pods(namespace).List(context.Background(), meta_v1.ListOptions{})
}