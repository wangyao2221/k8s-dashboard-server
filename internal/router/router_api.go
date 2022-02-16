package router

import "k8s-dashboard-server/internal/api/user"

func setApiRouter(r *resource) {
	// user
	userHandler := user.New(r.logger, r.db, r.cache)

	user := r.mux.Engine.Group("/api")
	{
		user.GET("/user", userHandler.List())
	}
}
