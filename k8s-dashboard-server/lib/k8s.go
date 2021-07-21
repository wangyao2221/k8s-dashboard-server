package lib

import (
	"k8s-dashboard-server/common"
	"k8s.io/client-go/kubernetes"
)

var (
	clientset *kubernetes.Clientset
)

func Clientset() (cs *kubernetes.Clientset, err error) {
	// 单例，加锁
	if clientset == nil {
		clientset, err = common.InitClient()
		if err != nil {
			return nil, err
		}
	}

	return clientset, nil
}
