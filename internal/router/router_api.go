package router

import (
	k8s "k8s-dashboard-server/internal/api/k8s"
)

func setApiRouter(r *resource) {
	k8sHandler := k8s.New(r.logger, r.k8s)
	k8s := r.mux.Group("/api/k8s")
	{
		k8s.GET("/namespaces", k8sHandler.Namespaces())
	}
}
