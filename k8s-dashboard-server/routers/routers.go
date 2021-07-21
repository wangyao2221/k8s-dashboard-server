package routers

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Register(r *gin.Engine)
}

type RootRouter struct {
	Engine *gin.Engine
}

func (rr RootRouter) Register()  {
	// v1
	v1 := rr.Engine.Group("/v1")
	// v1 api
	v1Api := v1.Group("/api")
	NewDeploymentRouter(v1Api).Register()
}
