package routers

import (
	"github.com/gin-gonic/gin"
	"k8s-dashboard-server/common"
	"k8s-dashboard-server/lib"
	"k8s-dashboard-server/services"
	core_v1 "k8s.io/api/core/v1"
	"net/http"
)

type DeploymentRouter struct {
	engine *gin.RouterGroup
	deploymentService *services.DeploymentService
}

func NewDeploymentRouter(engine *gin.RouterGroup) *DeploymentRouter {
	clientset, err := lib.Clientset()
	if err != nil {
		panic(err)
	}

	return &DeploymentRouter{
		engine:            engine,
		deploymentService: services.NewDeploymentService(clientset),
	}
}

func (dr *DeploymentRouter) Register()  {
	dr.engine.GET("/deployment", dr.Deployment)
}

func (dr *DeploymentRouter) Deployment(ctx *gin.Context) {
	var (
		podsList *core_v1.PodList
		err error
	)

	// 获取default命名空间下的所有POD
	if podsList, err = dr.deploymentService.List(common.DEFAULT_NAMESPACE); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"msg": "get pods from default namespace failed",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": podsList,
	})
}
