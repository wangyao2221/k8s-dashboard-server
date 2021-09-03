package common

import (
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// InitClient 初始化k8s客户端
func InitClient() (clientset *kubernetes.Clientset, err error) {
	var (
		restConf *rest.Config
	)

	if restConf, err = GetRestConf(); err != nil {
		return nil, err
	}

	// 生成clientset配置
	if clientset, err = kubernetes.NewForConfig(restConf); err != nil {
		return nil, err
	}

	return clientset, nil
}

// GetRestConf 获取k8s restful client配置
func GetRestConf() (restConf *rest.Config, err error) {
	var (
		kubeConfig []byte
	)

	// 读kubeConfig文件
	if kubeConfig, err = ioutil.ReadFile("./conf/docker-desktop-k8s.yaml"); err != nil {
		return nil, err
	}
	// 生成rest client配置
	if restConf, err = clientcmd.RESTConfigFromKubeConfig(kubeConfig); err != nil {
		return nil, err
	}
	return restConf, nil
}