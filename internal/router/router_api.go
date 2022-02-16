package router

import "k8s-dashboard-server/internal/api/helper"

func setApiRouter(r *resource) {
	// helper
	helperHandler := helper.New(r.logger, r.db, r.cache)

	helper := r.mux.Engine.Group("/helper")
	{
		helper.GET("/md5", helperHandler.Md5())
	}
}
