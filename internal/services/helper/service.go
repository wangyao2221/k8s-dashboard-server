package helper

import (
	"k8s-dashboard-server/internal/repository/mysql"
	"k8s-dashboard-server/internal/repository/redis"
)

type Service struct {
	db    mysql.Repo
	cache redis.Repo
}

func New(db mysql.Repo, cache redis.Repo) Service {
	return Service{
		db:    db,
		cache: cache,
	}
}
