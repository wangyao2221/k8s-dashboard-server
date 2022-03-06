package k8s

import (
	"go.uber.org/zap"
	"k8s-dashboard-server/internal/pkg/core"
	"k8s-dashboard-server/internal/repository/k8s"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Namespaces 获取集群namespaces
	// @tags K8S
	// @Router /k8s/namespaces
	Namespaces() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	k8s    k8s.Repo
}

func New(logger *zap.Logger, k8s k8s.Repo) Handler {
	return &handler{
		logger: logger,
		k8s:    k8s,
	}
}

func (h *handler) i() {}
