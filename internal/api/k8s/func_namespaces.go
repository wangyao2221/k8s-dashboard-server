package k8s

import (
	"context"
	"k8s-dashboard-server/internal/api"
	"k8s-dashboard-server/internal/code"
	"k8s-dashboard-server/internal/pkg/core"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

type namespacesData struct {
	Namespaces []string `json:"namespaces"`
}

func (h *handler) Namespaces() core.HandlerFunc {
	return func(ctx core.Context) {
		data := new(namespacesData)

		namespaces, err := h.k8s.GetClient().CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
		if err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.K8sError,
				code.Text(code.K8sError)).WithError(err),
			)
			return
		}

		data.Namespaces = []string{}
		for _, v := range namespaces.Items {
			data.Namespaces = append(data.Namespaces, v.Name)
		}

		ctx.Payload(api.Success(data))
	}
}
