package k8s

import (
	"context"
	"k8s-dashboard-server/internal/code"
	"k8s-dashboard-server/internal/pkg/core"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type namespacesResponse struct {
	Namespaces []string `json:"namespaces"`
}

func (h *handler) Namespaces() core.HandlerFunc {
	return func(ctx core.Context) {
		response := new(namespacesResponse)

		namespaces, err := h.k8s.GetClient().CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.K8sK8sError,
				code.Text(code.K8sK8sError)).WithError(err),
			)
			return
		}

		response.Namespaces = []string{}
		for _, v := range namespaces.Items {
			response.Namespaces = append(response.Namespaces, v.Name)
		}
		ctx.Payload(response)
	}
}
