package helper

import (
	"go.uber.org/zap"

	"k8s-dashboard-server/internal/repository/mysql"
	"k8s-dashboard-server/internal/repository/redis"
	"k8s-dashboard-server/internal/services/helper"
)

type Handler struct {
	logger        *zap.Logger
	db            mysql.Repo
	helperService helper.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return Handler{
		logger:        logger,
		db:            db,
		helperService: helper.New(db, cache),
	}
}
