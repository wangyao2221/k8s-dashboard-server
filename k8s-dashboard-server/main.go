package main

import (
	"github.com/gin-gonic/gin"
	"k8s-dashboard-server/routers"
)

func main() {
	r := gin.Default()

	router := routers.RootRouter{Engine: r}
	router.Register()
	r.Run(":8080")
}
