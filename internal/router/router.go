package router

import (
	"errors"
	"go.uber.org/zap"
	"k8s-dashboard-server/internal/repository/k8s"

	"k8s-dashboard-server/configs"
	"k8s-dashboard-server/internal/pkg/core"
	"k8s-dashboard-server/internal/repository/mysql"
	"k8s-dashboard-server/internal/repository/redis"
	"k8s-dashboard-server/internal/router/interceptor"
	"k8s-dashboard-server/pkg/file"
)

// 处理路由信息
type resource struct {
	mux          core.Mux
	logger       *zap.Logger
	db           mysql.Repo
	cache        redis.Repo
	k8s          k8s.Repo
	interceptors interceptor.Interceptor
	//cronServer   cron.Server
}

type Server struct {
	Mux   core.Mux
	Db    mysql.Repo
	Cache redis.Repo
}

// Server是衔接gin配置和router的中间桥梁
// 巧妙地关联了两者，又对两者进行了解耦(两者的代码互不依赖)
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
		//dbRepo, err := mysql.New()
		//if err != nil {
		//	logger.Fatal("new db err", zap.Error(err))
		//}
		//r.db = dbRepo
		//
		//// 初始化 Cache
		//cacheRepo, err := redis.New()
		//if err != nil {
		//	logger.Fatal("new cache err", zap.Error(err))
		//}
		//r.cache = cacheRepo

		// 初始化k8s
		k8sRepo, err := k8s.New()
		if err != nil {
			logger.Fatal("new cache err", zap.Error(err))
		}
		r.k8s = k8sRepo
	}

	// 封装一层engin的优势在这里，可以把一下默认参数的配置简化(例如这里的WithOption)
	mux, err := core.New(
		logger,
		core.WithEnableCors(),
	)

	if err != nil {
		panic(err)
	}

	r.mux = mux
	r.interceptors = interceptor.New(logger, r.cache, r.db)

	// 设置路由
	setApiRouter(r)

	s := new(Server)
	s.Mux = mux
	s.Db = r.db
	s.Cache = r.cache

	return s, nil
}
