package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	configs "k8s-dashboard-server/conifgs"
	"k8s-dashboard-server/internal/repository/mysql"
	"k8s-dashboard-server/internal/repository/redis"
	"k8s-dashboard-server/internal/router/interceptor"
	"k8s-dashboard-server/pkg/file"
)

// 处理路由信息
type resource struct {
	engine       *gin.Engine
	logger       *zap.Logger
	db           mysql.Repo
	cache        redis.Repo
	interceptors interceptor.Interceptor
	//cronServer   cron.Server
}

type Server struct {
	Engine *gin.Engine
	Db     mysql.Repo
	Cache  redis.Repo
}

func NewHTTPServer(logger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	r := new(resource)
	r.logger = logger

	openBrowserUri := configs.ProjectDomain + configs.ProjectPort
	_, ok := file.IsExists(configs.ProjectInstallMark)
	if !ok { // 未安装
		openBrowserUri += "/install"
	} else { // 已安装
		// 初始化 DB
		dbRepo, err := mysql.New()
		if err != nil {
			logger.Fatal("new db err", zap.Error(err))
		}
		r.db = dbRepo

		// 初始化 Cache
		cacheRepo, err := redis.New()
		if err != nil {
			logger.Fatal("new cache err", zap.Error(err))
		}
		r.cache = cacheRepo
	}

	s := new(Server)
	s.Db = r.db
	s.Cache = r.cache

	return s, nil
}
