package user

import (
	"go.uber.org/zap"

	"k8s-dashboard-server/internal/repository/mysql"
	"k8s-dashboard-server/internal/repository/redis"
	"k8s-dashboard-server/internal/services/user"
)

type Handler struct {
	logger      *zap.Logger
	db          mysql.Repo
	userService user.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return Handler{
		logger:      logger,
		db:          db,
		userService: user.New(db, cache),
	}
}
