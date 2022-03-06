package k8s

import (
	"k8s-dashboard-server/configs"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var _ Repo = (*k8sRepo)(nil)

type Repo interface {
	i()
	GetClient() *kubernetes.Clientset
}

type k8sRepo struct {
	client *kubernetes.Clientset
}

func New() (Repo, error) {
	cfg := configs.Get().K8S

	config, err := clientcmd.BuildConfigFromFlags("", cfg.ConfigPath)
	if err != nil {
		return nil, err
	}

	client, _ := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &k8sRepo{
		client: client,
	}, nil
}

func (k *k8sRepo) i() {}

func (k *k8sRepo) GetClient() *kubernetes.Clientset {
	return k.client
}
